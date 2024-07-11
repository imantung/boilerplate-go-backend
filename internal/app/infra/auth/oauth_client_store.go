package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/entity"
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
)

type OAuthClientStore struct {
	Repo entity.Oauth2ClientRepo
}

var _ oauth2.ClientStore = (*OAuthClientStore)(nil)
var _ = di.Provide(NewOAuthClientStore)

func NewOAuthClientStore(repo entity.Oauth2ClientRepo) *OAuthClientStore {
	return &OAuthClientStore{
		Repo: repo,
	}
}

func (o *OAuthClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	client, err := o.getClientByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Client{
		ID:     client.ClientID,
		Secret: client.Secret,
		Domain: client.Domain,
		UserID: strconv.Itoa(client.UserID),
	}, nil
}

func (o *OAuthClientStore) UserAuthorizationHandler(w http.ResponseWriter, req *http.Request) (userID string, err error) {
	params, _ := url.ParseQuery(req.URL.RawQuery)

	clientIDKey := "client_id"
	clientID := ""
	if len(params[clientIDKey]) > 0 {
		clientID = params[clientIDKey][0]
	}

	client, err := o.getClientByID(req.Context(), clientID)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(client.UserID), nil
}

func (o *OAuthClientStore) getClientByID(ctx context.Context, id string) (*entity.Oauth2Client, error) {
	clients, err := o.Repo.Select(ctx, repokit.Eq{"client_id": id})
	if err != nil {
		return nil, err
	}

	if len(clients) < 1 {
		return nil, fmt.Errorf("can't find client_id '%s'", id)
	}

	return clients[0], nil
}
