package app

import (
	"net/http"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/auth"
	"github.com/imantung/boilerplate-go-backend/internal/app/infra/logger"
	"github.com/imantung/boilerplate-go-backend/internal/app/service"
	"github.com/imantung/boilerplate-go-backend/internal/generated/oapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"

	_ "expvar"         // NOTE: enable `/debug/vars` endpoint
	_ "net/http/pprof" // NOTE: enable `/debug/pprof` endpoint
)

type Router struct {
	dig.In

	Health AppHealth

	Oauth              *auth.OAuthHandler
	BasicAuthValidator middleware.BasicAuthValidator

	service.EmployeeSvc
	service.ClockSvc
}

var _ oapi.StrictServerInterface = (*Router)(nil)

func (r *Router) SetRoute(e *echo.Echo) {

	e.HTTPErrorHandler = r.customErrorHandler

	e.Pre(
		middleware.RequestID(),
	)

	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Secure(),
		logger.HTTPRequest(),
	)

	// NOTE: register open-api endpoints
	oapi.RegisterHandlers(
		e.Group("api"),
		oapi.NewStrictHandler(r, []oapi.StrictMiddlewareFunc{
			r.Oauth.ValidateToken,
		}))

	e.Any("/oauth/authorize", r.Oauth.HandleAuthorizeRequest)
	e.Any("/oauth/token", r.Oauth.HandleTokenRequest)

	e.File("/swagger/api-spec.yaml", "api/api-spec.yaml")
	e.Static("/swagger/ui", "api/swagger-ui")

	basicAuth := middleware.BasicAuth(r.BasicAuthValidator)
	e.Any("/health", r.Health.Handle, basicAuth)
	e.GET("/debug/*/*", echo.WrapHandler(http.DefaultServeMux), basicAuth)
}

func (r *Router) customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "unknown error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message, _ = he.Message.(string)
	} else {
		message = err.Error()
	}
	c.Logger().Error(err)

	resp := oapi.Error{
		ErrorMessage: message,
	}
	if err := c.JSON(code, resp); err != nil {
		c.Logger().Error(err)
	}

}
