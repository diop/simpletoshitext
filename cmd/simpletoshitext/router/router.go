package router

import (
	"github.com/diop/simpletoshitext/cmd/simpletoshitext/groups"
	"github.com/diop/simpletoshitext/cmd/simpletoshitext/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Set all middlewares
	middlewares.SetMainMiddlewares(e)

	// Set main routes
	groups.MainGroup(e)

	return e
}
