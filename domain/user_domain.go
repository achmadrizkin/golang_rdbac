package domain

import "go-multirole/model"

type UserRepo interface {
	CreateUser(user model.User) (model.User, error)
	LoginUser(user model.User) (model.User, error)
	AssignRoleToUser(userId string, roleID string) error
	CheckUserPermission(userID string, permissionName string) (bool, error)
}

type UserUseCase interface {
	CreateUser(user model.User) (model.User, error)
	LoginUser(user model.User) (model.User, error)
	AssignRoleToUser(userId string, roleID string) error
	CheckUserPermission(userID string, permissionName string) (bool, error)
}
