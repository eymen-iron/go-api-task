package router

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	rows, err := DB.Query("SELECT ID as id, name, strftime('%Y-%m-%dT%H:%M:%SZ', start_date) as startDate, strftime('%Y-%m-%dT%H:%M:%SZ', end_date) as endDate, duration, durationUnit, color, externalId, status FROM construction_stages")
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Veritabanı hatası: %s", err.Error()),
		})
	}

	var constructionStages []ConstructionStage

	for rows.Next() {
		var startDate, endDate, durationUnit, color, externalID, status sql.NullString
		var duration sql.NullInt64
		var constructionStage ConstructionStage

		err := rows.Scan(&constructionStage.ID, &constructionStage.Name, &startDate, &endDate, &duration, &durationUnit, &color, &externalID, &status)
		if err != nil {
			return c.JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("Veritabanı hatası: %s", err.Error()),
			})
		}
		if color.Valid {
			constructionStage.Color = color.String
		}
		if externalID.Valid {
			constructionStage.ExternalID = externalID.String
		}
		if status.Valid {
			constructionStage.Status = status.String
		}

		if startDate.Valid {
			constructionStage.StartDate = startDate.String
		}

		if endDate.Valid {
			constructionStage.EndDate = endDate.String
		}

		if duration.Valid {
			constructionStage.Duration = int(duration.Int64)
		}
		if durationUnit.Valid {
			constructionStage.DurationUnit = durationUnit.String
		}

		constructionStages = append(constructionStages, constructionStage)
	}

	if err = rows.Err(); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Veritabanı hatası: %s", err.Error()),
		})
	}
	return c.JSON(constructionStages)
}
