package main

import (
	"log"
	"os"

	"hexa-fiber-gorm/internal/app/factory"
	"hexa-fiber-gorm/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbConn := db.NewConnection()

	app := fiber.New()

	// Bootstrap all modules
	modules := factory.BuildModules(dbConn)
	for _, m := range modules {
		m.RegisterRoutes(app)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
