package main

import (
	"math/rand"

	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	Id          uint
	Name        string
	Description string
	Price       int
}

func main() {

	db, err := gorm.Open(mysql.Open("root:root@/gomysqlfakedata"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	db.AutoMigrate(&Item{})
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/api/item/create", func(c *fiber.Ctx) error {

		for i := 0; i < 100; i++ {
			db.Create(&Item{
				Name:        faker.Word(),
				Description: faker.Paragraph(),
				Price:       rand.Intn(140) + 10,
			})

		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
		})
	})

	app.Get("/api/item/all", func(c *fiber.Ctx) error {

		var items []Item
		db.Find(&items)
		return c.Status(200).JSON(items)
	})

	app.Listen(":8000")
}
