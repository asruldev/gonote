package router

import (
	"gonote/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	routes.SetupNoteRoutes(api)
	routes.SetupUserRoutes(api)

}
