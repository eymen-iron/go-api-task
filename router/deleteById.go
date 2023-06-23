package router

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "The ID could not be converted to an integer.",
		})
	}

	constructionStage, err := DbGetSingle(idInt)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if constructionStage.Status == "DELETED" {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Construction stage already deleted.",
		})
	}

	err = DbDeleteById(idInt)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Construction stage successfully deleted.",
	})

}

func DbDeleteById(id int) (err error) {
	stmt, err := DB.Prepare(`
		UPDATE construction_stages
		SET 
			status = ?
		WHERE ID = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		"DELETED",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
