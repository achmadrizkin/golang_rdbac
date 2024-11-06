package usecase

import (
	"go-multirole/domain"
	"go-multirole/model"
)

type permissionUseCase struct {
	permissionRepo domain.PermissionRepo
}

func NewPermissionUseCase(permissionRepo domain.PermissionRepo) domain.PermissionUseCase {
	return &permissionUseCase{
		permissionRepo: permissionRepo,
	}
}

// CreateRole implements domain.RoleUseCase.
func (r *permissionUseCase) CreatePermission(role model.Permission) (model.Permission, error) {
	return r.permissionRepo.CreatePermission(role)
}
