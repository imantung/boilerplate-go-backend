package auth

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/rs/zerolog/log"
)

var _ = di.Provide(NewOAuthServer)

func NewOAuthServer(clientStore *OAuthClientStore) *server.Server {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore()) // token memory store
	manager.MapClientStorage(clientStore)

	// NOTE: skip URI validation if swagger-ui address is defferent with server address
	// manager.SetValidateURIHandler(func(baseURI, redirectURI string) error { return nil })

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.UserAuthorizationHandler = clientStore.UserAuthorizationHandler
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Err(err).Msg("Oauth Internal Error")
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Err(re.Error).Msg("Oauth Response Error")
	})

	return srv
}
