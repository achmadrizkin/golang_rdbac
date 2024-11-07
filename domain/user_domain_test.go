package domain

import (
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for UserRepo interface
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepo) LoginUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepo) AssignRoleToUser(userId string, roleID string) error {
	args := m.Called(userId, roleID)
	return args.Error(0)
}

func (m *MockUserRepo) CheckUserPermission(userID string, permissionName string) (bool, error) {
	args := m.Called(userID, permissionName)
	return args.Bool(0), args.Error(1)
}

// Mock for UserUseCase interface
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) CreateUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserUseCase) LoginUser(user model.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserUseCase) AssignRoleToUser(userId string, roleID string) error {
	args := m.Called(userId, roleID)
	return args.Error(0)
}

func (m *MockUserUseCase) CheckUserPermission(userID string, permissionName string) (bool, error) {
	args := m.Called(userID, permissionName)
	return args.Bool(0), args.Error(1)
}

// Unit Test for UserRepo interface
func TestUserRepo(t *testing.T) {
	mockRepo := new(MockUserRepo)

	// Test: Create User
	t.Run("Create User", func(t *testing.T) {
		user := model.User{ID: 1, Username: "john_doe", Password: "password123"}
		mockRepo.On("CreateUser", user).Return(user, nil)

		createdUser, err := mockRepo.CreateUser(user)

		assert.NoError(t, err)
		assert.Equal(t, user, createdUser)
		mockRepo.AssertExpectations(t)
	})

	// Test: Login User
	t.Run("Login User", func(t *testing.T) {
		user := model.User{Username: "john_doe", Password: "password123"}
		mockRepo.On("LoginUser", user).Return(user, nil)

		loggedInUser, err := mockRepo.LoginUser(user)

		assert.NoError(t, err)
		assert.Equal(t, user, loggedInUser)
		mockRepo.AssertExpectations(t)
	})

	// Test: Assign Role to User
	t.Run("Assign Role to User", func(t *testing.T) {
		userId := "1"
		roleID := "admin"
		mockRepo.On("AssignRoleToUser", userId, roleID).Return(nil)

		err := mockRepo.AssignRoleToUser(userId, roleID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	// Test: Check User Permission
	t.Run("Check User Permission", func(t *testing.T) {
		userId := "1"
		permissionName := "admin_access"
		mockRepo.On("CheckUserPermission", userId, permissionName).Return(true, nil)

		hasPermission, err := mockRepo.CheckUserPermission(userId, permissionName)

		assert.NoError(t, err)
		assert.True(t, hasPermission)
		mockRepo.AssertExpectations(t)
	})
}

// Unit Test for UserUseCase interface
func TestUserUseCase(t *testing.T) {
	mockUseCase := new(MockUserUseCase)

	// Test: Create User
	t.Run("Create User", func(t *testing.T) {
		user := model.User{ID: 1, Username: "john_doe", Password: "password123"}
		mockUseCase.On("CreateUser", user).Return(user, nil)

		createdUser, err := mockUseCase.CreateUser(user)

		assert.NoError(t, err)
		assert.Equal(t, user, createdUser)
		mockUseCase.AssertExpectations(t)
	})

	// Test: Login User
	t.Run("Login User", func(t *testing.T) {
		user := model.User{Username: "john_doe", Password: "password123"}
		mockUseCase.On("LoginUser", user).Return("1", nil)

		userId, err := mockUseCase.LoginUser(user)

		assert.NoError(t, err)
		assert.Equal(t, "1", userId)
		mockUseCase.AssertExpectations(t)
	})

	// Test: Assign Role to User
	t.Run("Assign Role to User", func(t *testing.T) {
		userId := "1"
		roleID := "admin"
		mockUseCase.On("AssignRoleToUser", userId, roleID).Return(nil)

		err := mockUseCase.AssignRoleToUser(userId, roleID)

		assert.NoError(t, err)
		mockUseCase.AssertExpectations(t)
	})

	// Test: Check User Permission
	t.Run("Check User Permission", func(t *testing.T) {
		userId := "1"
		permissionName := "admin_access"
		mockUseCase.On("CheckUserPermission", userId, permissionName).Return(true, nil)

		hasPermission, err := mockUseCase.CheckUserPermission(userId, permissionName)

		assert.NoError(t, err)
		assert.True(t, hasPermission)
		mockUseCase.AssertExpectations(t)
	})
}
