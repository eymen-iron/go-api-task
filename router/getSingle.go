package router

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "ID'yi tamsayıya çevirirken hata oluştu.",
		})
	}

	constructionStage, err := DbGetSingle(idInt)
	if err != nil {
		response := fiber.Map{
			"success": false,
			"message": err.Error(),
		}
		return c.JSON(response)
	}

	return c.JSON(constructionStage)
}

func DbGetSingle(id int) (ConstructionStage, error) {
	var constructionStage ConstructionStage
	var startDate, endDate, durationUnit, color, externalID, status sql.NullString
	var duration sql.NullInt64

	rows, err := DB.Prepare("SELECT ID as id, name, strftime('%Y-%m-%dT%H:%M:%SZ', start_date) as startDate, strftime('%Y-%m-%dT%H:%M:%SZ', end_date) as endDate, duration, durationUnit, color, externalId, status FROM construction_stages WHERE ID = ?")
	if err != nil {
		return constructionStage, err
	}

	err = rows.QueryRow(id).Scan(&constructionStage.ID, &constructionStage.Name, &startDate, &endDate, &duration, &durationUnit, &color, &externalID, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			return constructionStage, fmt.Errorf("ürün bulunamadı")
		}
		return constructionStage, err
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
	return constructionStage, nil
}
