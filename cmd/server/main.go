package main

import (
	"fmt"
	"toughleaf/cmd/server/routes"
	"toughleaf/config"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	config.ConnectDB()
	routes.Routes(app)

	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
