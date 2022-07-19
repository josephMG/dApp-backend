package services

import (
	"hardhat-backend/lib"
	"hardhat-backend/lib/loggers"
	"hardhat-backend/models"
	"hardhat-backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger          loggers.Logger
	repository      repository.UserRepository
	paginationScope *gorm.DB
}

// NewUserService creates a new userservice
func NewUserService(
	logger loggers.Logger,
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		logger:     logger,
		repository: userRepository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// PaginationScope
func (s UserService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) UserService {
	s.paginationScope = s.repository.WithTrx(s.repository.Scopes(scope)).DB
	return s
}

// GetOneUser gets one user
func (s UserService) GetOneUser(userID lib.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

// GetOneUser gets one user
func (s UserService) GetOneUserByWalletAddress(walletAddress string) (user models.User, newUser int64, err error) {
	result := s.repository.FirstOrCreate(&user, models.User{WalletAddress: walletAddress})
	return user, result.RowsAffected, result.Error
}

// GetAllUser get all the user
func (s UserService) GetAllUser() (response map[string]interface{}, err error) {
	var users []models.User
	var count int64

	err = s.repository.WithTrx(s.paginationScope).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"data": users, "count": count}, nil
}

// UpdateUser updates the user
func (s UserService) UpdateUser(user *models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(uuid lib.BinaryUUID) error {
	return s.repository.Where("id = ?", uuid).Delete(&models.User{}).Error
}

// DeleteUser deletes the user
func (s UserService) Create(user *models.User) error {
	return s.repository.Create(&user).Error
}
