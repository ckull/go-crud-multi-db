package types

import (
	"go-crud/config"
	"go-crud/pkg/factory"

	"github.com/labstack/echo/v4"
)

type (
	Server struct {
		App *echo.Echo
		Db  *factory.Database
		Cfg *config.Config
	}
)
