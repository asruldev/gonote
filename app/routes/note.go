package routes

import (
	"gonote/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/notes")
	// Create a Note
	note.Post("/", controllers.CreateNotes)
	// Read all Notes
	note.Get("/", controllers.GetNotes)
	// // Read one Note
	note.Get("/:noteId", controllers.GetNote)
	// // Update one Note
	note.Put("/:noteId", controllers.UpdateNote)
	// // Delete one Note
	note.Delete("/:noteId", controllers.DeleteNote)
}
