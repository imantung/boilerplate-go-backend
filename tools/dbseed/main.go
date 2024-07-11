package main

import (
	"context"
	"log"

	_ "github.com/imantung/boilerplate-go-backend/internal/app/infra/database"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
)

func main() {
	err := di.Invoke(dbseed)
	if err != nil {
		log.Fatal(err)
	}
}

func dbseed(oauthClientRepo entity.Oauth2ClientRepo) {
	ctx := context.Background()

	Must(oauthClientRepo.Insert(ctx, &entity.Oauth2Client{
		ClientID: "000000",
		UserID:   1,
		Secret:   "999999",
		Domain:   "http://localhost:1323",
	}))
}

func Must[T any](obj T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return obj
}
