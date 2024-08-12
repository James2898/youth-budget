package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello Universe")

	app := fiber.New()
	app.Get("/api/transaction", getTransactions)
	app.Post("/api/transaction", newTransaction)
	app.Patch("/api/transaction", updateTransaction)
	app.Delete("/api/transaction", deleteTransaction)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}
	app.Listen(":" + PORT)
}

func getTransactions(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"success": true})
}

func newTransaction(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"success": true})
}

func updateTransaction(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"success": true})
}

func deleteTransaction(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"success": true})
}
