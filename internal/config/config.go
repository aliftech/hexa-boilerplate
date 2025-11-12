// internal/config/config.go
package config

import "time"

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

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
