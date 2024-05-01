package repository

import (
	"context"
	"log"

	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{
		db: db,
	}
}

// FindByEmail implements model.UserRepository.
func (u *userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{}
	if err := u.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		log.Println(err)
		return model.User{}, err
	}

	return user, nil
}

// FindById implements model.UserRepository.
func (u *userRepository) FindById(ctx context.Context, userId string) (model.User, error) {
	user := model.User{}
	if err := u.db.WithContext(ctx).Where("id = ?", userId).Take(&user).Error; err != nil {
		log.Println(err)
		return model.User{}, err
	}

	return user, nil
}

// Insert implements model.UserRepository.
func (u *userRepository) Insert(ctx context.Context, user *model.User) error {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
