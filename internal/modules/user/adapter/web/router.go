// internal/modules/user/adapter/web/routes.go
package web

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterUserRoutes attaches user-related routes to the app
func RegisterUserRoutes(app *fiber.App, handler *Handler) {
	v1 := app.Group("/api/v1/users")

	v1.Post("/", handler.CreateUser)
	v1.Get("/:id", handler.GetUser)
	v1.Get("/", handler.GetAllUsers)
}
