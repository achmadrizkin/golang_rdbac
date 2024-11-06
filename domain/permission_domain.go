package domain

import "go-multirole/model"

type PermissionRepo interface {
	CreatePermission(permission model.Permission) (model.Permission, error)
}

type PermissionUseCase interface {
	CreatePermission(permission model.Permission) (model.Permission, error)
}
