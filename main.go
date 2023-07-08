package main

import (
	"zombie-golang/internal/domain/entities"
	database "zombie-golang/pkg/database/gorm"
	"zombie-golang/pkg/database/gorm/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := database.Bootstrap()

	app.Get("/", func(c *fiber.Ctx) error {
		var armor entities.Armor
		createdArmor := armor.New(entities.ArmorProps{Name: "Oi"})
		db.Create(&models.ArmorModel{
			Id:        createdArmor.Id,
			Name:      createdArmor.Name,
			CreatedAt: createdArmor.CreatedAt,
			UpdatedAt: createdArmor.UpdatedAt,
			DeletedAt: createdArmor.DeletedAt,
		})
		return c.JSON(createdArmor)
	})

	app.Listen(":3000")
}
