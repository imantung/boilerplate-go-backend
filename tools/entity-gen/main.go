package main

import (
	"log"

	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database" // NOTE: provide database constructor
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
)

func main() {
	if err := di.Invoke(generate); err != nil {
		log.Fatal(err)
	}
}
