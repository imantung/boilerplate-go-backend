package main

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
)

type (
	Table struct {
		PackageName string
		TableName   string
		Columns     []Column

		StructName string
	}
	Column struct {
		ColumnName string
		DataType   string
		IsNullable string

		FieldName     string
		FieldType     string
		IsPrimaryKey  bool
		IsAuditColumn bool
	}
)

var (
	PackageName  = "entity"
	TemplatePath = "tools/entity-gen/entity.go.tmpl"
	TargetDir    = "internal/generated/entity"
	MockDir      = "internal/generated/mock_entity"

	SkipTables   = []string{"schema_migrations"}
	PrimaryKeys  = []string{"id"}
	AuditColumns = []string{"deleted_at", "created_at", "updated_at"}
)

var pluralizer = pluralize.NewClient()

func main() {
	if err := di.Invoke(generate); err != nil {
		log.Fatal(err)
	}
}

func generate(db *sql.DB) error {
	tables, err := getTables(db)
	if err != nil {
		return err
	}

	os.MkdirAll(TargetDir, os.ModePerm)

	tmpl, err := template.ParseFiles(TemplatePath)
	if err != nil {
		return err
	}

	for _, table := range tables {
		targetPath := TargetDir + "/" + table.TableName + ".go"
		file, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer file.Close()

		if err := tmpl.Execute(file, table); err != nil {
			return err
		}
	}

	cmd := exec.Command("go", "run", "golang.org/x/tools/cmd/goimports@latest", "-w", TargetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	for _, table := range tables {
		cmd := exec.Command("go", "run", "github.com/golang/mock/mockgen@v1.6.0",
			"-destination", MockDir+"/mock_"+table.TableName+".go",
			"-source", TargetDir+"/"+table.TableName+".go",
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	return nil
}

func getTables(db *sql.DB) ([]Table, error) {
	tableNames, err := getTableNames(db)
	if err != nil {
		return nil, err
	}

	var tables []Table
	for _, tableName := range tableNames {
		if slices.Contains(SkipTables, tableName) {
			continue
		}
		columns, err := getColumns(db, tableName)
		if err != nil {
			return nil, err
		}

		tables = append(tables, Table{
			PackageName: PackageName,
			TableName:   tableName,
			StructName:  convertToStructName(tableName),
			Columns:     columns,
		})
	}

	return tables, nil
}

func getTableNames(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		return nil, err
	}

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func getColumns(db *sql.DB, table string) ([]Column, error) {
	rows, err := db.Query("SELECT column_name, data_type, is_nullable FROM information_schema.columns WHERE table_name = '" + table + "' ORDER BY ordinal_position")
	if err != nil {
		return nil, err
	}

	var columns []Column
	for rows.Next() {
		var column Column
		if err := rows.Scan(&column.ColumnName, &column.DataType, &column.IsNullable); err != nil {
			return nil, err
		}

		column.FieldName = convertToFieldName(column.ColumnName)                  // inject field name
		column.FieldType = convertToFieldType(column.DataType, column.IsNullable) // inject field type
		column.IsPrimaryKey = slices.Contains(PrimaryKeys, column.ColumnName)
		column.IsAuditColumn = slices.Contains(AuditColumns, column.ColumnName)

		columns = append(columns, column)
	}
	return columns, nil
}

func convertToStructName(tableName string) string {
	return strcase.ToCamel(pluralizer.Singular(tableName))
}

func convertToFieldName(colName string) string {
	if colName == "id" {
		return "ID"
	}

	colName = strcase.ToCamel(colName)
	if strings.HasSuffix(colName, "Id") {
		colName = colName[:len(colName)-2] + "ID"
	}
	return colName
}

func convertToFieldType(dataType string, isNullable string) string {
	fieldType := "UnknownType"

	if dataType == "integer" {
		fieldType = "int"
	}
	if dataType == "text" {
		fieldType = "string"
	}
	if strings.HasPrefix(dataType, "timestamp") {
		fieldType = "time.Time"
	}

	if strings.EqualFold(isNullable, "YES") {
		fieldType = "*" + fieldType
	}

	return fieldType
}
