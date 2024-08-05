package routes

import (
	"go-crud/modules/user/handler"
	repo "go-crud/modules/user/repository/mongodb"
	useCase "go-crud/modules/user/useCase/mongodb"
	"go-crud/server/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(s *types.Server) {
	dbConn := (*s.Db).GetConnection()
	dbType := s.Cfg.DBType

	var userRepo repo.UserRepository
	var userUsecase useCase.UserUsecase

	switch dbType {
	case "mongodb":
		mongoClient, ok := dbConn.(*mongo.Client)
		if !ok {
			log.Fatal("Failed to assert database connection as *mongo.Client")
		}
		userRepo = repo.NewUserRepository(mongoClient)
		userUsecase = useCase.NewUserUsecase(userRepo)
	case "postgres":
		// userRepo = postgres.NewUserRepository(db)
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}

	userHandler := handler.NewUserHandler(userUsecase)
	s.App.GET("/users", userHandler.GetUsers)
	s.App.POST("/users", userHandler.CreateUser)
	s.App.PUT("/users", userHandler.UpdateUser)
	s.App.DELETE("/users", userHandler.DeleteUser)

}
