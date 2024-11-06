package main

import (
	"go-multirole/config"
	"go-multirole/controller"
	"go-multirole/db"
	"go-multirole/repo"
	"go-multirole/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	db := db.InitDB(&loadConfig)
	router := gin.Default()

	userRepo := repo.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controller.NewUserController(userUseCase)

	roleRepo := repo.NewRoleRepository(db)
	roleUseCase := usecase.NewRoleUseCase(roleRepo)
	roleController := controller.NewRoleController(roleUseCase)

	permissionRepo := repo.NewPermissionRepository(db)
	permissionUseCase := usecase.NewPermissionUseCase(permissionRepo)
	permissionController := controller.NewPermissionController(permissionUseCase)

	// Define routes
	router.POST("/roles", roleController.CreateRole)
	router.POST("/permissions", permissionController.CreatePermission)

	router.POST("/users", userController.CreateUser)
	router.POST("/users/login", userController.LoginUser)

	router.GET("/users/:userID/roles/:roleID", userController.AssignRoleToUser)
	router.GET("/roles/:roleID/permissions/:permissionID", roleController.AssignPermissionToRole)
	router.GET("/users/:userID/permissions/:permissionName", userController.CheckUserPermission)

	router.GET("/users/temp", userController.GetUserTemp)

	router.Run(":9091")
}
