package cmd

import (
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server *ApiServer) setMiddleware() {
	server.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	server.echo.Use(middleware.Logger())
	server.echo.Use(middleware.Recover())
	server.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			os.Setenv("PROCESS_ID", uuid.New().String())
			return next(c)
		}
	})

	server.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if httpErr, ok := err.(*echo.HTTPError); ok {
			code = httpErr.Code
		}
		// resp := response.NewApiResponse(c.Path())
		// resp.Status.Code = code
		// resp.Status.Type = http.StatusText(code)
		c.JSON(code, err)
		// c.JSON(code, resp)
	}
}
