package cmd

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const API_SERVER_DEFAULT_PORT string = "8080"

type ApiServer struct {
	apiRunType string
	runtimeEnv string
	dbInstance string

	echo *echo.Echo
}

func (server *ApiServer) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = API_SERVER_DEFAULT_PORT
	}

	server.Start(port)
}

func (server *ApiServer) Start(port string) {
	errLoadEnv := server.loadEnv()
	if len(errLoadEnv) != 0 {
		for _, e := range errLoadEnv {
			log.Errorf(e)
		}
		return
	}

	server.echo = echo.New()
	server.setMiddleware()
}
