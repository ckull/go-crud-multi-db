package factory

import (
	"context"
	"go-crud/config"
	"go-crud/pkg/database"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type (
	Database interface {
		GetConnection() interface{}
		Disconnect() interface{}
	}

	MongoDB struct {
		Client *mongo.Client
	}

	PostgresDB struct {
		DB *gorm.DB
	}
)

func (db *MongoDB) GetConnection() interface{} {
	return db.Client
}

func (db *PostgresDB) GetConnection() interface{} {
	return db.DB
}

func (db *MongoDB) Disconnect() interface{} {
	return db.Client.Disconnect(context.Background())
}

func (db *PostgresDB) Disconnect() interface{} {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func NewDatabase(ctx context.Context, cfg *config.Config) Database {
	switch cfg.DBType {
	case "mongodb":
		{
			client := database.MongoConn(ctx, cfg)

			return &MongoDB{
				Client: client,
			}
		}

	case "postgres":
		{
			db := database.PostgresConn(ctx, cfg)

			return &PostgresDB{
				DB: db,
			}
		}

	default:
		{
			log.Fatal("Unsupported database type")
			return nil
		}

	}

}
