package factory

import (
	"hexa-fiber-gorm/internal/core"
	"hexa-fiber-gorm/pkg/db"
)

var modules []func(*db.Connection) core.Module

func RegisterModule(creator func(*db.Connection) core.Module) {
	modules = append(modules, creator)
}

func BuildModules(dbConn *db.Connection) []core.Module {
	var built []core.Module
	for _, create := range modules {
		built = append(built, create(dbConn))
	}
	return built
}
