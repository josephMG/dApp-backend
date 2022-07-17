package repository

import (
	"hardhat-backend/infrastructure"
	"hardhat-backend/lib/loggers"

	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	infrastructure.Database
	logger loggers.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db infrastructure.Database, logger loggers.Logger) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx delegate transaction from user repository
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}
