package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConfig matches internal/config.DBConfig (or embed it)
type DBConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	TimeZone        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	Env             string
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.TimeZone,
	)
}

type GORM struct {
	db *gorm.DB
}

// New accepts *DBConfig from internal/config
func New(config *DBConfig) (*GORM, error) {
	// ... (rest unchanged)
	var logLevel logger.LogLevel
	switch config.Env {
	case "development", "staging":
		logLevel = logger.Info
	default:
		logLevel = logger.Error
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	db, err := gorm.Open(mysql.Open(config.DSN()), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL DB: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("MySQL unreachable: %w", err)
	}

	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	}

	return &GORM{db: db}, nil
}

func (g *GORM) DB() *gorm.DB { return g.db }
func (g *GORM) Close() error {
	sqlDB, _ := g.db.DB()
	return sqlDB.Close()
}
