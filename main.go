package main

import (
	"gonote/database"
	"gonote/router"

	"log"

	_ "gonote/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title AsrulDev API
// @version 1.0
// @description Server api for blog
// @termsOfService https://asrul.dev/terms/

// @contact.name API Support
// @contact.url https://asrul.dev/support
// @contact.email talkasrul@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5555
// @BasePath /
// @schemes http
func main() {
	// Start a new fiber app
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	app.Get("/swagger/*", swagger.Handler) // default
	// Routes
	app.Get("/", HealthCheck)

	// Start Server
	if err := app.Listen(":5555"); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
