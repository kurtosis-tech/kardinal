package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	kardinal_manager_server_rest_server "kardinal.kontrol/kardinal-manager/api/http_rest/server"
	"net"
)

const (
	pathToApiGroup         = "/api"
	restAPIPortAddr uint16 = 8080
	restAPIHostIP   string = "0.0.0.0"
)

var (
	defaultCORSOrigins = []string{"*"}
	defaultCORSHeaders = []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}
)

func CreateAndStartRestAPIServer() error {
	logrus.Info("Running REST API server...")

	// This is how you set up a basic Echo router
	echoRouter := echo.New()
	echoApiRouter := echoRouter.Group(pathToApiGroup)
	echoApiRouter.Use(middleware.Logger())

	// CORS configuration
	echoApiRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: defaultCORSOrigins,
		AllowHeaders: defaultCORSHeaders,
	}))

	server := NewServer()

	kardinal_manager_server_rest_server.RegisterHandlers(echoApiRouter, kardinal_manager_server_rest_server.NewStrictHandler(server, nil))

	return echoRouter.Start(net.JoinHostPort(restAPIHostIP, fmt.Sprint(restAPIPortAddr)))
}
