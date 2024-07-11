package auth

import (
	"net/http"
	"slices"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
	strict "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

type (
	OAuthHandler struct {
		Server *server.Server
	}
)

var _ = di.Provide(NewOAuthHandler)

var HeaderXUserID = "X-User-Id"

func NewOAuthHandler(srv *server.Server) *OAuthHandler {
	return &OAuthHandler{Server: srv}
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

func (o *OAuthHandler) HandleAuthorizeRequest(c echo.Context) error {
	return o.Server.HandleAuthorizeRequest(c.Response().Writer, c.Request())
}

func (o *OAuthHandler) HandleTokenRequest(c echo.Context) error {
	return o.Server.HandleTokenRequest(c.Response().Writer, c.Request())
}
