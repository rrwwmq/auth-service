package main

import (
	"fmt"
	"os"

	"github.com/rrwwmq/auth-service/internal/repository/postgres"
	"github.com/rrwwmq/auth-service/internal/service"
)

func main() {
	fmt.Println("Hello, World!")
	db, err := postgres.New(os.Getenv("CONN_STRING"))
	defer db.Close()
	if err != nil {
		fmt.Println("unable to connect to database")
	}

	userRepo := postgres.NewUserRepo(db)
	authService := service.NewAuthService(userRepo)
	err = authService.Register("gazirovka228@gmail.com", "123123")
	if err != nil {
		return
	}

}
