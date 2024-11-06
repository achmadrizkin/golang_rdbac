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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := d.userUseCase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userResponse)
}

func (d *UserController) LoginUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := d.userUseCase.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "User login success",
	})
}

func (d *UserController) AssignRoleToUser(c *gin.Context) {
	userID := c.Param("userID")
	roleID := c.Param("roleID")

	err := d.userUseCase.AssignRoleToUser(userID, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable assign user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permission assigned to role"})
}

func (d *UserController) CheckUserPermission(c *gin.Context) {
	userID := c.Param("userID")
	permissionName := c.Param("permissionName")

	hasPermission, err := d.userUseCase.CheckUserPermission(userID, permissionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"has_permission": hasPermission})
}

func (d *UserController) GetUserTemp(c *gin.Context) {
	userID := c.Param("userID")
	permissionName := "read"

	_, err := d.userUseCase.CheckUserPermission(userID, permissionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found, or doesnt have access: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tempUser": "usernameTemp"})
}
