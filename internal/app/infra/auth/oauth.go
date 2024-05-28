package auth

import (
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

const (
	OAuth2Token = "OAuth2.Token"
)

type (
	OAuthHandler struct {
		Server *server.Server
	}
)

var _ = di.Provide(NewHandler)

func NewHandler() *OAuthHandler {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore()) // token memory store

	// // NOTE: by passed validate URI due to different uri between swagger-ui address and server address
	// manager.SetValidateURIHandler(func(baseURI, redirectURI string) error {
	// 	return nil
	// })

	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{ // TODO: create API set client
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:1323",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	handler := &OAuthHandler{
		Server: srv,
	}

	srv.UserAuthorizationHandler = handler.UserAuthorizationHandler
	srv.SetInternalErrorHandler(handler.InternalErrorHandler)
	srv.SetResponseErrorHandler(handler.ResponseErrorHandler)

	return handler
}

func (o *OAuthHandler) ValidateTokenMW() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			token, err := o.Server.ValidationBearerToken(req)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			c.Set(OAuth2Token, token)
			return next(c)
		}
	}
}

func (o *OAuthHandler) InternalErrorHandler(err error) (re *errors.Response) {
	log.Println("Internal Error:", err.Error()) // TODO: change the logger
	return
}

func (o *OAuthHandler) ResponseErrorHandler(re *errors.Response) {
	log.Println("Response Error:", re.Error.Error()) // TODO: change the logger
}

func (o *OAuthHandler) UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	return "000000", nil
}

func (o *OAuthHandler) HandleAuthorizeRequest(c echo.Context) error {
	req := c.Request()
	w := c.Response().Writer
	return o.Server.HandleAuthorizeRequest(w, req)
}

func (o *OAuthHandler) HandleTokenRequest(c echo.Context) error {
	req := c.Request()
	w := c.Response().Writer
	return o.Server.HandleTokenRequest(w, req)
}
