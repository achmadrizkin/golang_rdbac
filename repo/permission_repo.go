package repo

import (
	"go-multirole/domain"
	"go-multirole/model"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) domain.PermissionRepo {
	return &permissionRepository{
		db: db,
	}
}

// CreatePermission implements domain.PermissionRepo.
func (p *permissionRepository) CreatePermission(permission model.Permission) (model.Permission, error) {
	if err := p.db.Create(&permission).Error; err != nil {
		return permission, err
	}
	return permission, nil
}
