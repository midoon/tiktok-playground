package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/dto"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primary_key;column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email;uniqueIndex"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		id := uuid.New().String()
		u.ID = id
	}

	return nil
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindById(ctx context.Context, userId string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) error
	Login(ctx context.Context, req *dto.LoginRequest) (dto.UserData, error)
}
