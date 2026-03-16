package main

import (
	"fmt"
	"os"

	"github.com/rrwwmq/auth-service/internal/repository/postgres"
	"github.com/rrwwmq/auth-service/internal/service"
	"github.com/rrwwmq/auth-service/internal/transport/rest"
)

func main() {
	db, err := postgres.New(os.Getenv("CONN_STRING"))
	if err != nil {
		fmt.Println("failed to connect to database", err)
		return
	}
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)
	authService := service.NewAuthService(userRepo)
	handler := rest.NewHandler(authService)
	router := handler.InitRoutes()
	fmt.Println("start server on port :8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Println("failed to connect server")
		return
	}
}
