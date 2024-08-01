package server

import (
	"context"
	"go-crud/config"
	"go-crud/modules/user/routes"
	"go-crud/pkg/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Server struct {
		app *echo.Echo
		db  *factory.Database
		cfg *config.Config
	}
)

func Start(ctx context.Context, cfg *config.Config, db *factory.Database) {

	s := &Server{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}

	// CORS
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	routes.UserRoute(&s)

}
