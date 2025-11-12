package web

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, handler *Handler) {
	v1 := app.Group("/api/v1/users")
	v1.Post("/", handler.CreateUser)
	v1.Get("/:id", handler.GetUser)
	v1.Get("/", handler.GetAllUsers)
}
