package usecase

import (
	"go-multirole/domain"
	"go-multirole/model"
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
func (u *userUseCase) LoginUser(user model.User) (model.User, error) {
	return u.userRepo.LoginUser(user)
}
