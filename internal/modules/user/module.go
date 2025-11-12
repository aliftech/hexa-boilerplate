package user

import (
	"hexa-fiber-gorm/internal/app/factory"
	"hexa-fiber-gorm/internal/core"
	"hexa-fiber-gorm/internal/modules/user/adapter/persistence"
	"hexa-fiber-gorm/internal/modules/user/adapter/web"
	"hexa-fiber-gorm/internal/modules/user/usecase"
	"hexa-fiber-gorm/pkg/db"

	"github.com/gofiber/fiber/v2"
)

// userModule implements core.Module
type userModule struct {
	db *db.GORM
}

func (m *userModule) RegisterRoutes(app *fiber.App) {
	// Build dependency chain: repo → usecase → handler
	repo := persistence.NewUserRepository(m.db)
	usecase := usecase.NewUserUsecase(repo)
	handler := web.NewHandler(usecase)

	web.RegisterUserRoutes(app, handler)
}

// Factory function
func newUserModule(db *db.GORM) core.Module {
	return &userModule{db: db}
}

// Auto-register this module on package init
func init() {
	factory.RegisterModule(newUserModule)
}
