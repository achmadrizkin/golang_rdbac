package repo

import (
	"go-multirole/domain"
	"go-multirole/model"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepo {
	return &roleRepository{
		db: db,
	}
}

// CreateRole implements domain.RoleRepo.
func (r *roleRepository) CreateRole(role model.Role) (model.Role, error) {
	if err := r.db.Create(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

// AssignPermissionToRole implements domain.RoleRepo.
func (r *roleRepository) AssignPermissionToRole(roleID string, permissionID string) error {
	var role model.Role
	var permission model.Permission

	if err := r.db.First(&role, roleID).Error; err != nil {
		return err
	}
	if err := r.db.First(&permission, permissionID).Error; err != nil {
		return err
	}

	if err := r.db.Model(&role).Association("Permissions").Append(&permission); err != nil {
		return err
	}

	return nil
}
