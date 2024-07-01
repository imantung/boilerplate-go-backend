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

var HeaderXUserID = "X-User-Id"

func NewHandler() *OAuthHandler {
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{ // TODO: create API set client
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:1323",
	})

	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore()) // token memory store
	manager.MapClientStorage(clientStore)

	// NOTE: skip URI validation if swagger-ui address is defferent with server address
	// manager.SetValidateURIHandler(func(baseURI, redirectURI string) error { return nil })

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	handler := &OAuthHandler{Server: srv}
	srv.UserAuthorizationHandler = handler.UserAuthorizationHandler
	srv.SetInternalErrorHandler(handler.InternalErrorHandler)
	srv.SetResponseErrorHandler(handler.ResponseErrorHandler)

	return handler
}

func (o *OAuthHandler) UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	return "000000", nil
}

func (o *OAuthHandler) ValidateToken(next strict.StrictEchoHandlerFunc, operationID string) strict.StrictEchoHandlerFunc {
	return func(c echo.Context, req interface{}) (resp interface{}, err error) {
		httpReq := c.Request()
		token, err := o.Server.ValidationBearerToken(httpReq)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		apiScopes, _ := c.Get(oapi.OAuth2Scopes).([]string)
		if !validateScope(token, apiScopes) {
			return nil, echo.NewHTTPError(http.StatusForbidden, "user scopes not match")
		}

		c.Request().Header.Set(HeaderXUserID, token.GetUserID()) // NOTE: inject new request header
		return next(c, req)
	}
}

func validateScope(token oauth2.TokenInfo, apiScopes []string) bool {
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

func (o *OAuthHandler) HandleAuthorizeRequest(c echo.Context) error {
	return o.Server.HandleAuthorizeRequest(c.Response().Writer, c.Request())
}

func (o *OAuthHandler) HandleTokenRequest(c echo.Context) error {
	return o.Server.HandleTokenRequest(c.Response().Writer, c.Request())
}
