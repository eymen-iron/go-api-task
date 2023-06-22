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
			"message": err,
		})
	}

	data := make(map[string]string)
	message := map[string]interface{}
	for _, v := range formArr {
		if c.FormValue(v) != "" {
			data[v] = c.FormValue(v)
		}
	}

	for k, v := range data {
		switch k {
		case "name":
			message[k] = validator.GetValidate(k, v)
		case "startDate":
			message[k] = validator.GetValidate(k, v)
		case "endDate":
			message[k] = validator.GetValidate(k, v)
		case "duration":
			message[k] = validator.GetValidate(k, v)
		case "durationUnit":
			message[k] = validator.GetValidate(k, v)
		case "color":
			message[k] = validator.GetValidate(k, v)
		case "externalId":
			message[k] = validator.GetValidate(k, v)
		case "status":
			message[k] = validator.GetValidate(k, v)
		}
	}

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "ID'yi tamsayıya çevirirken hata oluştu.",
		})
	}

	return c.JSON(fiber.Map{
		"id":   idInt,
		"Data": data,
		"Name": constructionStage,
	})

}
