package main

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/fatih/structs"
	"github.com/iancoleman/strcase"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra"
)

func main() {
	cfg := &infra.Config{}
	prefix := infra.ConfigPrefix
	writeDotenv(os.Stdout, prefix, cfg)
}

func writeDotenv(w io.Writer, prefix string, val interface{}) {
	fields := structs.Fields(val)

	for _, field := range fields {
		key := getKey(prefix, field)
		if field.Kind() == reflect.Struct {
			writeDotenv(w, key, field.Value())
		} else {
			value := getValue(field)
			fmt.Fprintln(w, key+"="+value)
		}
	}
}

func getKey(prefix string, field *structs.Field) string {
	key := field.Tag("envconfig")
	if key == "" {
		key = strcase.ToScreamingSnake(field.Name())
	}
	if prefix != "" {
		key = prefix + "_" + key
	}
	return key
}

func getValue(field *structs.Field) string {
	return field.Tag("default")
}
