package factory

import (
	"hexa-fiber-gorm/internal/core"
	"hexa-fiber-gorm/pkg/db"

	"github.com/gofiber/fiber/v2"
)

// moduleCreator builds a module from shared dependencies
type moduleCreator func(*db.GORM) core.Module

var modules []moduleCreator

// RegisterModule adds a module factory to the registry
func RegisterModule(creator moduleCreator) {
	modules = append(modules, creator)
}

// BuildModules constructs all modules and returns a bootstrapable group
func BuildModules(gormDB *db.GORM) *ModuleGroup {
	var built []core.Module
	for _, create := range modules {
		built = append(built, create(gormDB))
	}
	return &ModuleGroup{modules: built}
}

// ModuleGroup holds all initialized modules and can register their routes
type ModuleGroup struct {
	modules []core.Module
}

// RegisterRoutes registers all module routes to the Fiber app
func (mg *ModuleGroup) RegisterRoutes(app *fiber.App) {
	for _, module := range mg.modules {
		module.RegisterRoutes(app)
	}
}
