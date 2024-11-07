package controller

import (
	"bytes"
	"errors"
	"go-multirole/model"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the UserUseCase interface
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

func (m *MockUserUseCase) AssignRoleToUser(userID string, roleID string) error {
	args := m.Called(userID, roleID)
	return args.Error(0)
}

func (m *MockUserUseCase) CheckUserPermission(userID string, permissionName string) (bool, error) {
	args := m.Called(userID, permissionName)
	return args.Bool(0), args.Error(1)
}

// Test for CreateUser
func TestCreateUser(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	userController := NewUserController(mockUseCase)

	t.Run("Create user successfully", func(t *testing.T) {
		// Define a mock user that matches the request body and expected response
		mockUser := model.User{ID: 1, Username: "john_doe", Password: "password123"}

		// Mock the return of CreateUser with the mock user and no error
		mockUseCase.On("CreateUser", mock.AnythingOfType("model.User")).Return(mockUser, nil)

		// Create a test HTTP request with a JSON body that matches the user model
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/user", bytes.NewBufferString(`{"username":"john_doe", "password":"password123"}`))

		// Call the CreateUser function
		userController.CreateUser(c)

		// Assert the response status is OK and the message is as expected
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Created user success")

		// Verify mock expectations
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Create user with invalid JSON", func(t *testing.T) {
		// Create a test HTTP request with invalid JSON
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/user", bytes.NewBufferString(`{"username":"john_doe"`)) // invalid JSON

		// Call the CreateUser function
		userController.CreateUser(c)

		// Assert the response status is BadRequest
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "unexpected EOF")
	})
}

// Test for LoginUser
func TestLoginUser(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	userController := NewUserController(mockUseCase)

	t.Run("Login user successfully", func(t *testing.T) {
		// Define a mock user and token
		mockUser := model.User{Username: "john_doe", Password: "password123"}
		mockToken := "mock_token"
		mockUseCase.On("LoginUser", mockUser).Return(mockToken, nil)

		// Create a test HTTP request and recorder
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`{"username":"john_doe", "password":"password123"}`))

		// Call the LoginUser function
		userController.LoginUser(c)

		// Assert the response status and body
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Login Success")
		assert.Contains(t, w.Body.String(), mockToken)

		// Verify mock expectations
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Login user with error", func(t *testing.T) {
		// Define a mock user and simulate an error during login
		mockUser := model.User{Username: "john_doe", Password: "wrong_password"}
		mockUseCase.On("LoginUser", mockUser).Return("", errors.New("invalid credentials"))

		// Create a test HTTP request and recorder
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`{"username":"john_doe", "password":"wrong_password"}`))

		// Call the LoginUser function
		userController.LoginUser(c)

		// Assert the response status and error message
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "invalid credentials")

		// Verify mock expectations
		mockUseCase.AssertExpectations(t)
	})
}

// Test for AssignRoleToUser
func TestAssignRoleToUser(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	userController := NewUserController(mockUseCase)

	t.Run("Assign role to user successfully", func(t *testing.T) {
		mockUseCase.On("AssignRoleToUser", "1", "2").Return(nil)

		// Create a test HTTP request and recorder
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "userID", Value: "1"}, gin.Param{Key: "roleID", Value: "2"}}
		c.Request, _ = http.NewRequest(http.MethodPost, "/assign-role", nil)

		// Call the AssignRoleToUser function
		userController.AssignRoleToUser(c)

		// Assert the response status and message
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Permission assigned to role to user")

		// Verify mock expectations
		mockUseCase.AssertExpectations(t)
	})
}

// Test for CheckUserPermission
func TestCheckUserPermission(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	userController := NewUserController(mockUseCase)

	t.Run("Check user permission successfully", func(t *testing.T) {
		mockUseCase.On("CheckUserPermission", "1", "read").Return(true, nil)

		// Create a test HTTP request and recorder
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "userID", Value: "1"}, gin.Param{Key: "permissionName", Value: "read"}}
		c.Request, _ = http.NewRequest(http.MethodGet, "/check-permission", nil)

		// Call the CheckUserPermission function
		userController.CheckUserPermission(c)

		// Assert the response status and body
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"has_permission":true`)

		// Verify mock expectations
		mockUseCase.AssertExpectations(t)
	})
}
