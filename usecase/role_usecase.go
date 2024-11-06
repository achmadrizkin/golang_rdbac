package usecase

import (
	"go-multirole/domain"
	"go-multirole/model"
)

type roleUseCase struct {
	roleRepo domain.RoleRepo
}

func NewRoleUseCase(roleRepo domain.RoleRepo) domain.RoleUseCase {
	return &roleUseCase{
		roleRepo: roleRepo,
	}
}

// CreateRole implements domain.RoleUseCase.
func (r *roleUseCase) CreateRole(role model.Role) (model.Role, error) {
	return r.roleRepo.CreateRole(role)
}

// AssignPermissionToRole implements domain.RoleUseCase.
func (r *roleUseCase) AssignPermissionToRole(roleID string, permissionID string) error {
	return r.roleRepo.AssignPermissionToRole(roleID, permissionID)
}
