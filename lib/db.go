package lib

import (
	"fmt"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		User:     GetEnv("POSTGRES_USER", "postgres"),
		Password: GetEnv("POSTGRES_PASSWORD", "postgres"),
		Host:     GetEnv("POSTGRES_HOST", "localhost"),
		Port:     GetEnv("POSTGRES_PORT", "5432"),
		Name:     GetEnv("POSTGRES_DB", "learning_platform"),
	}
}

func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}
