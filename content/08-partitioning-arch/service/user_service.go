package service

import (
	"context"

	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/dto"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/helper"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/model"
)

type userService struct {
	userRepository model.UserRepository
}

func NewUserService(userRepo model.UserRepository) model.UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// Login implements model.UserService.
func (u *userService) Login(ctx context.Context, req *dto.LoginRequest) (dto.UserData, error) {
	user, err := u.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.UserData{}, err
	}

	if user == (model.User{}) {
		return dto.UserData{}, err
	}

	if user.Password != req.Password {
		return dto.UserData{}, err
	}

	return dto.UserData{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// Register implements model.UserService.
func (u *userService) Register(ctx context.Context, req *dto.RegisterRequest) error {
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := u.userRepository.Insert(ctx, &user); err != nil {
		return helper.ErrRegistration
	}
	return nil
}
