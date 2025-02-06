package postgres

import (
	"fmt"

	"github.com/kyosyun/mmf-back/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg env.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", cfg.Host, cfg.User, cfg.Pass, cfg.Name, cfg.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{}) //nolint:wrapcheck
}
