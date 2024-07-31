package routes

import (
	"go-crud/modules/user/handler"
	"go-crud/modules/user/repository/mongodb"
	"go-crud/modules/user/useCase/mongodb"
	"go-crud/server"
	"log"
)

func UserRoute(s *server.Server) {
	db := s.db.GetConnection()
	dbType := s.cfg.DBType

	var userRepo any
	var userUsecase any

	switch dbType {
	case "mongodb":
		userRepo = mongodb.NewUserRepository(db)
		userUsecase = mongodb.NewUserUsecase(userRepo)
	case "postgres":
		// userRepo = postgres.NewUserRepository(db)
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}

	userHandler := handler.NewUserHandler(userUsecase)
	s.app.GET("/users", userHandler.GetUsers)
	s.app.POST("/users", userHandler.CreateUser)
	s.app.PUT("/users", userHandler.UpdateUser)
	s.app.DELETE("/users", userHandler.DeleteUser)

}
