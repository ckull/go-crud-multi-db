package database

import (
	"context"
	"go-crud/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConn(ctx context.Context, cfg *config.Config) *gorm.DB {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db, err := gorm.Open(postgres.Open(cfg.Db.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
