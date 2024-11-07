package usecase

import (
	"errors"
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the PermissionRepo interface
type MockPermissionRepo struct {
	mock.Mock
}

func (m *MockPermissionRepo) CreatePermission(permission model.Permission) (model.Permission, error) {
	args := m.Called(permission)
	return args.Get(0).(model.Permission), args.Error(1)
}

func TestCreatePermission(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockPermissionRepo)

	// Define the test input and expected output
	testPermission := model.Permission{
		ID:   1,
		Name: "Test Permission",
	}

	// Set up expectations: mock the CreatePermission method
	mockRepo.On("CreatePermission", testPermission).Return(testPermission, nil)

	// Create the UseCase with the mocked repository
	useCase := NewPermissionUseCase(mockRepo)

	// Call the method under test
	result, err := useCase.CreatePermission(testPermission)

	// Assert the expectations
	assert.NoError(t, err)
	assert.Equal(t, testPermission, result)

	// Assert that the CreatePermission method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestCreatePermission_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockPermissionRepo)

	// Define the test input
	testPermission := model.Permission{
		ID:   1,
		Name: "Test Permission",
	}

	// Set up expectations: simulate an error returned by CreatePermission
	mockRepo.On("CreatePermission", testPermission).Return(model.Permission{}, errors.New("failed to create permission"))

	// Create the UseCase with the mocked repository
	useCase := NewPermissionUseCase(mockRepo)

	// Call the method under test
	result, err := useCase.CreatePermission(testPermission)

	// Assert that an error occurred
	assert.Error(t, err)
	assert.Equal(t, model.Permission{}, result)

	// Assert that the CreatePermission method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}
