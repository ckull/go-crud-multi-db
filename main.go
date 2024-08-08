package main

import (
	"context"
	"go-crud/config"
	"go-crud/pkg/factory"
	"go-crud/server"
	"log"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:8080
// @BasePath /api/v1
func main() {

	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	db := factory.NewDatabase(ctx, cfg)

	server.Start(ctx, cfg, &db)

}
