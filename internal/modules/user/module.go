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

type userModule struct {
	dbConn *db.Connection
}

func (m *userModule) RegisterRoutes(app *fiber.App) {
	repo := persistence.NewUserRepository(m.dbConn)

	// ✅ Build usecase from repo
	usecase := usecase.NewUserUsecase(repo) // returns port.UserUsecase

	// ✅ Pass usecase to handler
	handler := web.NewHandler(usecase) // handler expects UserUsecase

	web.RegisterUserRoutes(app, handler)
}

func newUserModule(dbConn *db.Connection) core.Module {
	return &userModule{dbConn: dbConn}
}

func init() {
	factory.RegisterModule(func(dbConn *db.Connection) core.Module {
		return newUserModule(dbConn)
	})
}
