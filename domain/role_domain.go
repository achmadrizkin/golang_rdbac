package domain

import "go-multirole/model"

type RoleRepo interface {
	CreateRole(role model.Role) (model.Role, error)
	AssignPermissionToRole(roleID string, permissionID string) error
}

type RoleUseCase interface {
	CreateRole(role model.Role) (model.Role, error)
	AssignPermissionToRole(roleID string, permissionID string) error
}
