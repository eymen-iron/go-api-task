package router

import (
	"github.com/eymen-iron/go-api-task/validator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func UpdateByID(c *fiber.Ctx) error {
	formArr := []string{"name", "startDate", "endDate", "duration", "durationUnit", "color", "externalId", "status"}
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
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	data := make(map[string]string)
	var messageVals []validator.ValidatorMessage
	var messages []string
	for _, v := range formArr {
		if c.FormValue(v) != "" {
			data[v] = c.FormValue(v)
		}
	}

	for k, v := range data {
		switch k {
		case "name":
			messageVals = append(messageVals, validator.Name(v))
		case "startDate":
			messageVals = append(messageVals, validator.StartDate(v))
		case "endDate":
			messageVals = append(messageVals, validator.EndDate(v, constructionStage.StartDate))
		case "durationUnit":
			messageVals = append(messageVals, validator.DurationUnit(v))
		case "color":
			messageVals = append(messageVals, validator.Color(v))
		case "externalId":
			messageVals = append(messageVals, validator.ExternalId(v))
		case "status":
			messageVals = append(messageVals, validator.Status(v))
		}
	}

	for _, v := range messageVals {
		if v.Error == true {
			messages = append(messages, v.Message)
		}
	}

	if len(messages) > 0 {
		return c.JSON(fiber.Map{
			"success": false,
			"message": messages,
		})
	}

	for k, v := range data {
		switch k {
		case "name":
			constructionStage.Name = v
		case "startDate":
			constructionStage.StartDate = v
		case "endDate":
			constructionStage.EndDate = v
		case "durationUnit":
			constructionStage.DurationUnit = v
		case "color":
			constructionStage.Color = v
		case "externalId":
			constructionStage.ExternalID = v
		case "status":
			constructionStage.Status = v
		}
	}

	constructionStage.Duration = validator.CalculateDuration(constructionStage.StartDate, constructionStage.EndDate, constructionStage.DurationUnit)

	err = DbUpdateSingle(constructionStage)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Güncelleme başarılı.",
	})

}

func DbUpdateSingle(constructionStage ConstructionStage) error {
	stmt, err := DB.Prepare(`
		UPDATE construction_stages
		SET
			name = ?,
			start_date = ?,
			end_date = ?,
			duration = ?,
			durationUnit = ?,
			color = ?,
			externalId = ?,
			status = ?
		WHERE ID = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		constructionStage.Name,
		constructionStage.StartDate,
		constructionStage.EndDate,
		constructionStage.Duration,
		constructionStage.DurationUnit,
		constructionStage.Color,
		constructionStage.ExternalID,
		constructionStage.Status,
		constructionStage.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
