package controller

import (
	"go-multirole/domain"
	"go-multirole/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleUseCase domain.RoleUseCase
}

func NewRoleController(roleUseCase domain.RoleUseCase) *RoleController {
	return &RoleController{roleUseCase}
}

func (d *RoleController) CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleResponse, err := d.roleUseCase.CreateRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to create role: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Created role success",
		Data:       roleResponse,
	})
}

func (d *RoleController) AssignPermissionToRole(c *gin.Context) {
	roleID := c.Param("roleID")
	permissionID := c.Param("permissionID")

	err := d.roleUseCase.AssignPermissionToRole(roleID, permissionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to create role: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Permission assigned to role",
	})
}
