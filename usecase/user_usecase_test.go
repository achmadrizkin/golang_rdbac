package usecase

import (
	"errors"
	"go-multirole/config"
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the UserRepo interface
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user model.User) (model.User, error) {
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

func (m *MockUserRepo) LoginUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

// Mocking utils functions
func MockVerifyPassword(expectedPassword, actualPassword string) bool {
	return expectedPassword == actualPassword
}

func MockGenerateToken(expiresIn int, userID uint, secret string) (string, error) {
	if userID == 0 {
		return "", errors.New("invalid user id")
	}
	return "mock_token", nil
}

// Mocking config.LoadConfig function
func MockLoadConfig(path string) (*config.Config, error) {
	return &config.Config{
		TokenExpiresIn: 3600,
		TokenSecret:    "mock_secret",
	}, nil
}

func TestCreateUser(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepo)

	// Define the test input and expected output
	testUser := model.User{
		ID:       1,
		Username: "testuser",
		Password: "password",
	}

	// Set up expectations: mock the CreateUser method
	mockRepo.On("CreateUser", testUser).Return(testUser, nil)

	// Create the UseCase with the mocked repository
	useCase := NewUserUseCase(mockRepo)

	// Call the method under test
	result, err := useCase.CreateUser(testUser)

	// Assert the expectations
	assert.NoError(t, err)
	assert.Equal(t, testUser, result)

	// Assert that the CreateUser method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestAssignRoleToUser(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepo)

	// Define the test input
	userID := "1"
	roleID := "101"

	// Set up expectations: mock the AssignRoleToUser method
	mockRepo.On("AssignRoleToUser", userID, roleID).Return(nil)

	// Create the UseCase with the mocked repository
	useCase := NewUserUseCase(mockRepo)

	// Call the method under test
	err := useCase.AssignRoleToUser(userID, roleID)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the AssignRoleToUser method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestCheckUserPermission(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockUserRepo)

	// Define the test input
	userID := "1"
	permissionName := "admin"

	// Set up expectations: mock the CheckUserPermission method
	mockRepo.On("CheckUserPermission", userID, permissionName).Return(true, nil)

	// Create the UseCase with the mocked repository
	useCase := NewUserUseCase(mockRepo)

	// Call the method under test
	result, err := useCase.CheckUserPermission(userID, permissionName)

	// Assert the expectations
	assert.NoError(t, err)
	assert.True(t, result)

	// Assert that the CheckUserPermission method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}
