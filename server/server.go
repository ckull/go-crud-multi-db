package server

import (
	"context"
	"go-crud/config"
	"go-crud/modules/user/routes"
	"go-crud/server/types"

	"go-crud/pkg/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "go-crud/docs"
)

func Start(ctx context.Context, cfg *config.Config, db *factory.Database) {
	s := &types.Server{
		App: echo.New(),
		Db:  db,
		Cfg: cfg,
	}

	// CORS
	s.App.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	routes.UserRoute(s)

	s.App.GET("/swagger/*", echoSwagger.WrapHandler)
	s.App.Logger.Fatal(s.App.Start(":8080"))
}
