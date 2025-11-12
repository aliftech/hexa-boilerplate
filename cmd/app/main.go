package main

import (
	"log"

	"hexa-fiber-gorm/internal/app/factory"
	"hexa-fiber-gorm/internal/config"
	"hexa-fiber-gorm/pkg/db"

	// 🔌 Force-load all modules (triggers init → auto-registration)
	_ "hexa-fiber-gorm/internal/modules/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	// 2. Initialize DB (infrastructure)
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

	// 3. Start app
	app := fiber.New()
	factory.BuildModules(gormDB).RegisterRoutes(app)

	log.Printf("🚀 Server running on http://localhost:%s (env: %s)", cfg.Server.Port, cfg.Server.Env)
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
