package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"mini-wallet-exercise/Lib"
	"mini-wallet-exercise/Middleware"
	"mini-wallet-exercise/Model/Database"
	"mini-wallet-exercise/Route"
	"os"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	Lib.InitAll()

	// run migration
	err := Lib.DB.AutoMigrate(
		&Database.User{},
		&Database.Wallet{},
		&Database.ReferenceLog{},
	)
	if err != nil {
		panic(err)
	}

	Middleware.Init(e)
	Route.Init(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
