package repo

import (
	"go-multirole/domain"
	"go-multirole/model"

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
	if err := d.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
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
	hasPermission := false

	if err := d.db.Preload("Roles.Permissions").First(&user, userID).Error; err != nil {
		return false, err
	}

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
