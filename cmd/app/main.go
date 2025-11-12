package main

import (
	"log"

	"hexa-fiber-gorm/internal/app/factory"
	"hexa-fiber-gorm/internal/config"
	"hexa-fiber-gorm/internal/i18n"
	"hexa-fiber-gorm/internal/middleware"
	"hexa-fiber-gorm/pkg/db"

	_ "hexa-fiber-gorm/internal/modules/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load locales
	i18n.LoadLocales("locales")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	// Initialize DB
	gormDB, err := db.New(&db.DBConfig{
		Host:            cfg.DB.Host,
		Port:            cfg.DB.Port,
		User:            cfg.DB.User,
		Password:        cfg.DB.Password,
		DBName:          cfg.DB.DBName,
		TimeZone:        cfg.DB.TimeZone,
		MaxOpenConns:    cfg.DB.MaxOpenConns,
		MaxIdleConns:    cfg.DB.MaxIdleConns,
		ConnMaxLifetime: cfg.DB.ConnMaxLifetime,
		Env:             cfg.DB.Env,
	})
	if err != nil {
		log.Fatal("❌ Database init failed:", err)
	}
	defer gormDB.Close()

	// Setup Fiber
	app := fiber.New()

	// Middleware
	app.Use(middleware.NewCORSMiddleware())

	// Register modules
	_ = "hexa-fiber-gorm/internal/modules/user" // force import
	factory.BuildModules(gormDB).RegisterRoutes(app)

	port := cfg.Server.Port
	if port == "" {
		port = "3000"
	}
	log.Printf("🚀 Server running on http://localhost:%s (env: %s)", port, cfg.Server.Env)
	log.Fatal(app.Listen(":" + port))
}
