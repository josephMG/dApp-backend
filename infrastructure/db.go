package infrastructure

import (
	"fmt"
	"hardhat-backend/config"
	"hardhat-backend/lib/loggers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
	dsn string
}

// NewDatabase creates a new database instance
func NewDatabase(logger loggers.Logger, env *config.Env) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	logger.Info("Database connection established")

	return Database{
		DB:  db,
		dsn: url,
	}
}
