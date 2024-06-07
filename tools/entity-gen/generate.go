package main

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
)

type (
	Definition struct {
		PackageName string
		Tables      []Table
	}
	Table struct {
		TableName string
		Columns   []Column

		StructName string
	}
	Column struct {
		ColumnName string
		DataType   string
		IsNullable string

		FieldName string
		FieldType string
	}
)

var (
	PackageName  = "entity"
	TemplatePath = "tools/entity-gen/entity.go.tmpl"
	TargetPath   = "internal/generated/entity/entity.go"
	SkipTables   = map[string]any{
		"schema_migrations": nil,
	}
)

func main() {
	if err := di.Invoke(generate); err != nil {
		log.Fatal(err)
	}
}

func generate(db *sql.DB) error {
	def, err := getDefinition(db)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(TemplatePath)
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Dir(TargetPath), os.ModePerm)
	file, err := os.Create(TargetPath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, def); err != nil {
		return err
	}

	file.Close()

	return exec.
		Command("go", "run", "golang.org/x/tools/cmd/goimports@latest", "-w", TargetPath).
		Run()
}

func getDefinition(db *sql.DB) (*Definition, error) {
	tableNames, err := getTableNames(db)
	if err != nil {
		return nil, err
	}

	var tables []Table
	for _, tableName := range tableNames {
		if _, ok := SkipTables[tableName]; ok {
			continue
		}
		columns, err := getColumns(db, tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, Table{
			TableName:  tableName,
			StructName: strcase.ToCamel(tableName),
			Columns:    columns,
		})
	}

	return &Definition{
		PackageName: PackageName,
		Tables:      tables,
	}, nil
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

		columns = append(columns, column)
	}
	return columns, nil
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
