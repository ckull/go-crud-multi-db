package mongodb

import (
	"context"
	"go-crud/config"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dbInstance *mongo.Client
	once       sync.Once
)

func MongoConn(ctx context.Context, cfg *config.Config) *mongo.Client {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		dbInstance, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.MongoDBURI))
		if err != nil {
			log.Fatalf("Error: Connect to database error: %s", err.Error())
		}

		if err := dbInstance.Ping(ctx, readpref.Primary()); err != nil {
			log.Fatalf("Error: Pinging to database error: %s", err.Error())
		}
	})

	return dbInstance
}
