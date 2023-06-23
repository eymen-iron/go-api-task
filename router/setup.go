package router

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

var DB *sql.DB

type ConstructionStage struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	Duration     int    `json:"duration"`
	DurationUnit string `json:"durationUnit"`
	Color        string `json:"color"`
	ExternalID   string `json:"externalId"`
	Status       string `json:"status"`
}

func SetupRoutes(app *fiber.App, db *sql.DB) {
	DB = db
	app.Get("/", GetAll)
	app.Get("/:id", GetSingle)
	app.Patch("/:id", UpdateByID)
	app.Delete("/:id", DeleteByID)
}
