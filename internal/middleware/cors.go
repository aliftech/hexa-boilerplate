package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// NewCORSMiddleware returns configured CORS middleware
func NewCORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:5173,http://127.0.0.1:5173", // adjust for your frontend
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	})
}
