package main

import (
	"context"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/Mockturnal/voting-app-backend/models"
	"github.com/Mockturnal/voting-app-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	database.Init()
	if err := database.Ping(context.TODO()); err != nil {
		panic(err)
	}

	database.Migrate(&models.User{}, &models.Poll{})

	router := routes.Init()

	router.Logger.Fatal(router.Start(":5000"))
}
