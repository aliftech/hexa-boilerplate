package core

import "github.com/gofiber/fiber/v2"

// Module defines the contract every module must implement
type Module interface {
	RegisterRoutes(app *fiber.App)
}
