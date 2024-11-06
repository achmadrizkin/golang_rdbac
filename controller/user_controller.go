package controller

import (
	"go-multirole/domain"
	"go-multirole/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase domain.UserUseCase
}

func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{userUseCase}
}

func (d *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	userResponse, err := d.userUseCase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Created user success",
		Data:       userResponse,
	})
}

func (d *UserController) LoginUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	token, err := d.userUseCase.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Login Success",
		Data:       token,
	})
}

func (d *UserController) AssignRoleToUser(c *gin.Context) {
	userID := c.Param("userID")
	roleID := c.Param("roleID")

	err := d.userUseCase.AssignRoleToUser(userID, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusOK,
			Message:    "Unable assign user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Permission assigned to role to user",
	})
}

func (d *UserController) CheckUserPermission(c *gin.Context) {
	userID := c.Param("userID")
	permissionName := c.Param("permissionName")

	hasPermission, err := d.userUseCase.CheckUserPermission(userID, permissionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "User not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"has_permission": hasPermission})
}

func (d *UserController) GetUserTemp(c *gin.Context) {
	userID := c.MustGet("currentUserId").(string)
	permissionName := "read"

	has_permission, err := d.userUseCase.CheckUserPermission(userID, permissionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusOK,
			Message:    "User not found, or doesnt have access: " + err.Error(),
		})
		return
	}

	if has_permission {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    "Temp user and role",
		})
	} else {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "User doesnt have access",
		})
	}
}
