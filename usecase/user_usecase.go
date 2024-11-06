package usecase

import (
	"errors"
	"go-multirole/config"
	"go-multirole/domain"
	"go-multirole/model"
	"go-multirole/utils"
)

type userUseCase struct {
	userRepo domain.UserRepo
}

func NewUserUseCase(userRepo domain.UserRepo) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) CreateUser(user model.User) (model.User, error) {
	return u.userRepo.CreateUser(user)
}

// AssignRoleToUser implements domain.UserUseCase.
func (u *userUseCase) AssignRoleToUser(userId string, roleID string) error {
	return u.userRepo.AssignRoleToUser(userId, roleID)
}

// CheckUserPermission implements domain.UserUseCase.
func (u *userUseCase) CheckUserPermission(userID string, permissionName string) (bool, error) {
	return u.userRepo.CheckUserPermission(userID, permissionName)
}

// LoginUser implements domain.UserUseCase.
func (u *userUseCase) LoginUser(user model.User) (string, error) {
	dbUser, err := u.userRepo.LoginUser(user)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(dbUser.Password, user.Password) {
		return "", errors.New("incorrect password")
	}

	config, err := config.LoadConfig(".")
	if err != nil {
		return "", nil
	}

	token, err := utils.GenerateToken(config.TokenExpiresIn, dbUser.ID, config.TokenSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
