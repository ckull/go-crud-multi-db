package routes

import (
	mongoHandler "go-crud/modules/user/handler/mongodb"
	postgresHandler "go-crud/modules/user/handler/postgres"
	mongoRepo "go-crud/modules/user/repository/mongodb"
	postgresRepo "go-crud/modules/user/repository/postgres"
	mongoUseCase "go-crud/modules/user/useCase/mongodb"
	postgresUseCase "go-crud/modules/user/useCase/postgres"
	"go-crud/server/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func UserRoute(s *types.Server) {
	dbConn := (*s.Db).GetConnection()

	dbType := s.Cfg.DBType

	var userRepo interface{}
	var userUsecase interface{}
	var userHandler interface{}

	switch dbType {
	case "mongodb":
		mongoClient, ok := dbConn.(*mongo.Client)
		if !ok {
			log.Fatal("Failed to assert database connection as *mongo.Client")
		}
		userRepo = mongoRepo.NewUserRepository(mongoClient)
		userUsecase = mongoUseCase.NewUserUsecase(userRepo.(mongoRepo.UserRepository))
		userHandler = mongoHandler.NewUserHandler(userUsecase.(mongoUseCase.UserUsecase))

		s.App.GET("/users", userHandler.(mongoHandler.UserHandler).GetUsers)
		s.App.POST("/users", userHandler.(mongoHandler.UserHandler).CreateUser)
		s.App.PUT("/users", userHandler.(mongoHandler.UserHandler).UpdateUser)
		s.App.DELETE("/users", userHandler.(mongoHandler.UserHandler).DeleteUser)
	case "postgres":
		postgresClient, ok := dbConn.(*gorm.DB)
		if !ok {
			log.Fatal("Failed to assert database connection as *mongo.Client")
		}
		userRepo = postgresRepo.NewUserRepository(postgresClient)
		userUsecase = postgresUseCase.NewUserUsecase(userRepo.(postgresRepo.UserRepository))
		userHandler = postgresHandler.NewUserHandler(userUsecase.(postgresUseCase.UserUsecase))

		s.App.GET("/users", userHandler.(postgresHandler.UserHandler).GetUsers)
		s.App.POST("/users", userHandler.(postgresHandler.UserHandler).CreateUser)
		s.App.PUT("/users", userHandler.(postgresHandler.UserHandler).UpdateUser)
		s.App.DELETE("/users", userHandler.(postgresHandler.UserHandler).DeleteUser)

	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}

}
