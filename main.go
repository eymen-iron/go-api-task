package main

import (
	"log"

	"github.com/eymen-iron/go-api-task/db"
	"github.com/eymen-iron/go-api-task/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database := db.Database{}
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	router.SetupRoutes(app, db)

	err = app.Listen(":3005")
	if err != nil {
		log.Fatal(err)
	}
}
