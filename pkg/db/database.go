// Package db provides MySQL database connection using GORM.
package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConfig holds MySQL connection settings.
type DBConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	TimeZone        string // e.g., "Local" or "Asia%2FJakarta"
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	Env             string // "development", "staging", "production"
}

// DSN returns the MySQL Data Source Name string.
func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.TimeZone,
	)
}

// GORM wraps the GORM database connection.
type GORM struct {
	db *gorm.DB
}

// New creates and returns a new MySQL database connection using GORM.
func New(config *DBConfig) (*GORM, error) {
	var logLevel logger.LogLevel
	switch config.Env {
	case "development", "staging":
		logLevel = logger.Info
	default: // production or unknown
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
		return nil, fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("MySQL server is unreachable: %w", err)
	}

	// Configure connection pool
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

// DB returns the underlying *gorm.DB instance.
func (g *GORM) DB() *gorm.DB {
	return g.db
}

// Close closes the database connection.
func (g *GORM) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
