package postgres

import (
	"context"
	"go-crud/config"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func PostgresConn(ctx context.Context, cfg *config.Config) *gorm.DB {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		dbInstance, err := gorm.Open(postgres.Open(cfg.Db.PostgresDSN), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		sqlDB, err := dbInstance.DB()
		if err != nil {
			log.Fatal(err)
		}

		dbErr := sqlDB.PingContext(ctx)
		if dbErr != nil {
			log.Fatal("Error: Pinging to database error: ", dbErr)
		}

	})

	return dbInstance

}
