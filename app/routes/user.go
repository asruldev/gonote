package routes

import (
	"gonote/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	// Register an user
	user.Post("/register", controllers.RegisterUser)
	user.Post("/login", controllers.LoginUser)
	user.Get("/me", controllers.User)
	user.Post("/logout", controllers.LogoutUser)
	user.Post("/change-avatar", controllers.ChangeAvatar)
	user.Delete("/remove-avatar", controllers.RemoveAvatar)
}
