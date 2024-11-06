package controller

import (
	"go-multirole/domain"
	"go-multirole/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	permissionUseCase domain.PermissionUseCase
}

func NewPermissionController(permissionUseCase domain.PermissionUseCase) *PermissionController {
	return &PermissionController{permissionUseCase}
}

func (d *PermissionController) CreatePermission(c *gin.Context) {
	var permission model.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permissionResponse, err := d.permissionUseCase.CreatePermission(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to create permission: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Created permission success",
		Data:       permissionResponse,
	})
}
