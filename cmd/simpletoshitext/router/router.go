package router

import (
	"projects/echo-server-template/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Set all middlewares
	middlewares.SetMainMiddlewares(e)

	// Set main routes
	api.MainGroup(e)

	return e
}
