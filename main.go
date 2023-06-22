package main

import (
	"log"

	"github.com/eymen-iron/go-api-task/db"
	"github.com/eymen-iron/go-api-task/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Veritabanını başlat
	database := db.Database{}
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fiber uygulamasını oluştur
	app := fiber.New()

	// Router'ı yapılandır ve rotaları ayarla
	router.SetupRoutes(app, db)

	// Sunucuyu başlat
	err = app.Listen(":3005")
	if err != nil {
		log.Fatal(err)
	}
}
