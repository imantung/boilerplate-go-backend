package auth

import (
	"net/http"
	"slices"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
	strict "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	"github.com/rs/zerolog/log"
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

	handler := &OAuthHandler{Server: srv}
	srv.UserAuthorizationHandler = handler.UserAuthorizationHandler
	srv.SetInternalErrorHandler(handler.InternalErrorHandler)
	srv.SetResponseErrorHandler(handler.ResponseErrorHandler)

	return handler
}

func (o *OAuthHandler) ValidateToken(next strict.StrictEchoHandlerFunc, operationID string) strict.StrictEchoHandlerFunc {
	return func(c echo.Context, req interface{}) (resp interface{}, err error) {
		httpReq := c.Request()
		token, err := o.Server.ValidationBearerToken(httpReq)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		apiScopes, _ := c.Get(oapi.OAuth2Scopes).([]string)
		if !o.validateScope(token, apiScopes) {
			return nil, echo.NewHTTPError(http.StatusForbidden, "user scopes not match")
		}

		return next(c, req)
	}
}

func (o *OAuthHandler) validateScope(token oauth2.TokenInfo, apiScopes []string) bool {
	userScopes := strings.Split(token.GetScope(), " ")
	for _, scope := range apiScopes {
		if !slices.Contains(userScopes, scope) {
			return false
		}
	}
	return true
}

func (o *OAuthHandler) InternalErrorHandler(err error) (re *errors.Response) {
	log.Err(err).Msg("Oauth Internal Error")
	return
}

func (o *OAuthHandler) ResponseErrorHandler(re *errors.Response) {
	log.Err(re.Error).Msg("Oauth Response Error")
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
