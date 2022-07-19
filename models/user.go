package models

import (
	"database/sql"
	"hardhat-backend/lib"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	lib.ModelBase
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Birthday      *time.Time     `json:"time"`
	WalletAddress string         `json:"walletAddress" gorm:"uniqueIndex"`
	MemberNumber  sql.NullString `json:"memberNumber"`
	//	ProfilePic   lib.SignedURL  `json:"profile_pic"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BeforeCreate run this before creating user
func (t *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}
