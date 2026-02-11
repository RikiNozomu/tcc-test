package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(36);primary_key;"`
	Username  string `gorm:"unique"`
	Password  string // hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

type UserCreate struct {
	Username string
	Password string // hashed password
}

type UserRepository interface {
	Create(user UserCreate) (*User, error)
	GetOne(id string) (*User, error)
	GetByUsername(username string) (*User, error)
}
