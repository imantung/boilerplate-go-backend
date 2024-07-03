package main

import (
	"log"
	"os"
	"path"
	"reflect"
	"slices"
	"sort"
	"strings"
	"text/template"

	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/imantung/boilerplate-go-backend/pkg/cmdkit"
)

type (
	Data struct {
		PackageName string
		Converters  []*Converter
	}
	Converter struct {
		FnName     string
		FromType   string
		ToType     string
		ToTypeInit string
		FieldPairs []FieldPair
	}
	Struct struct {
		Name   string
		Fields []string
	}
	FieldPair struct {
		Left  string
		Right string
	}
)

var (
	TemplatePath  = "tools/converter-gen/converter.go.tmpl"
	TargetPath    = "internal/generated/converter/converter.go"
	PackageName   = "converter"
	SkippedFields = []string{"CreatedAt", "UpdatedAt", "DeletedAt"}
)

func main() {
	data := &Data{
		PackageName: PackageName,
		Converters: []*Converter{
			createConverter("ConvertToEmployeeOApi", entity.Employee{}, oapi.Employee{}),
			createConverter("ConvertToEmployeeEntity", oapi.Employee{}, entity.Employee{}),
			createConverter("ConvertToClockHistoryOApi", entity.EmployeeClockHistory{}, oapi.EmployeeClockHistory{}),
		},
	}

	tmpl, err := template.ParseFiles(TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	targetDir := path.Dir(TargetPath)
	os.Mkdir(targetDir, os.ModePerm)

	file, err := os.Create(TargetPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		log.Fatal(err)
	}

	cmdkit.GoImports(targetDir)
}

func createConverter(fnName string, from, to any) *Converter {
	fromStruct := createStruct(from)
	toStruct := createStruct(to)

	if len(fromStruct.Fields) != len(toStruct.Fields) {
		log.Fatalf("can't create converter function from %s to %s", fromStruct.Name, toStruct.Name)
	}

	fromType := fromStruct.Name
	toType := toStruct.Name
	toTypeInit := toType
	if strings.HasPrefix(toType, "*") {
		toTypeInit = "&" + toType[1:]
	}

	var fieldPairs []FieldPair
	for i, field := range toStruct.Fields {
		fieldPairs = append(fieldPairs, FieldPair{Left: field, Right: fromStruct.Fields[i]})
	}

	return &Converter{
		FnName:     fnName,
		FromType:   fromType,
		ToType:     toType,
		ToTypeInit: toTypeInit,
		FieldPairs: fieldPairs,
	}
}

func createStruct(i any) *Struct {
	val := reflect.ValueOf(i)
	typ := val.Type()

	var elem reflect.Type
	if typ.Kind() == reflect.Struct {
		elem = typ
	} else if typ.Kind() == reflect.Pointer {
		elem = typ.Elem()
	} else {
		panic("wrong type")
	}

	var fields []string
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i).Name
		if !slices.Contains(SkippedFields, field) {
			fields = append(fields, field)
		}
	}

	sort.Strings(fields)

	return &Struct{
		Name:   typ.String(),
		Fields: fields,
	}
}
