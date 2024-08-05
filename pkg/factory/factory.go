package factory

import (
	"context"
	"go-crud/config"
	"go-crud/modules/user/model/mongodb"
	"go-crud/modules/user/model/postgres"
	"go-crud/pkg/database"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type (
	Database interface {
		GetConnection() interface{}
		Disconnect() interface{}
		GetModel(modelType string) interface{}
	}

	MongoDB struct {
		Client *mongo.Client
		Models map[string]interface{}
	}

	PostgresDB struct {
		DB     *gorm.DB
		Models map[string]interface{}
	}
)

func (db *MongoDB) GetConnection() interface{} {
	return db.Client
}

func (db *PostgresDB) GetConnection() interface{} {
	return db.DB
}

func (db *MongoDB) GetModel(modelType string) interface{} {
	return db.Models[modelType]
}

func (db *PostgresDB) GetModel(modelType string) interface{} {
	return db.Models[modelType]
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
				Models: map[string]interface{}{
					"user": &mongodb.User{},
					// Add other MongoDB Models here
				},
			}
		}

	case "postgres":
		{
			db := database.PostgresConn(ctx, cfg)

			return &PostgresDB{
				DB: db,
				Models: map[string]interface{}{
					"user": &postgres.User{},
					// Add other MongoDB Models here
				},
			}
		}

	default:
		{
			log.Fatal("Unsupported database type")
			return nil
		}

	}

}
