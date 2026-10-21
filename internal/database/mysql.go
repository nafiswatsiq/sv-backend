package database

import (
	"fmt"

	"sharingvision/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) (*gorm.DB, error) {
	if cfg.DBConnection != "" && cfg.DBConnection != "mysql" {
		return nil, fmt.Errorf("unsupported database connection: %s", cfg.DBConnection)
	}

	host := cfg.DBHost
	if host == "" {
		host = "localhost"
	}

	baseDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		host,
		cfg.DBPort,
	)

	baseDB, err := gorm.Open(mysql.Open(baseDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := baseDB.Exec("CREATE DATABASE IF NOT EXISTS `" + cfg.DBName + "` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci").Error; err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		host,
		cfg.DBPort,
		cfg.DBName,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
