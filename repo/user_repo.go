package repo

import (
	"errors"
	"fmt"
	"go-multirole/domain"
	"go-multirole/model"
	"go-multirole/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepo {
	return &userRepository{
		db: db,
	}
}

// CreateUser implements domain.UserRepo.
func (d *userRepository) CreateUser(user model.User) (model.User, error) {
	user.Password, _ = utils.HashPassword(user.Password)
	if err := d.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// LoginUser checks credentials and returns the authenticated user with roles.
func (d *userRepository) LoginUser(inputUser model.User) (model.User, error) {
	var dbUser model.User

	if err := d.db.Where("username = ?", inputUser.Username).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, fmt.Errorf("database error: %w", err)
	}

	return dbUser, nil
}

// AssignRoleToUser implements domain.UserRepo.
func (d *userRepository) AssignRoleToUser(userId string, roleID string) error {
	var user model.User
	var role model.Role

	if err := d.db.First(&user, userId).Error; err != nil {
		return err
	}
	if err := d.db.First(&role, roleID).Error; err != nil {
		return err
	}
	if err := d.db.Model(&user).Association("Roles").Append(&role); err != nil {
		return err
	}

	return nil
}

// CheckUserPermission implements domain.UserRepo.
func (d *userRepository) CheckUserPermission(userID string, permissionName string) (bool, error) {
	var user model.User

	if err := d.db.Preload("Roles.Permissions").First(&user, userID).Error; err != nil {
		return false, err
	}

	hasPermission := false
	for _, role := range user.Roles {
		for _, perm := range role.Permissions {
			if perm.Name == permissionName {
				hasPermission = true
				break
			}
		}
		if hasPermission {
			break
		}
	}

	return hasPermission, nil
}
