package usecase

import (
	"errors"
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the RoleRepo interface
type MockRoleRepo struct {
	mock.Mock
}

func (m *MockRoleRepo) CreateRole(role model.Role) (model.Role, error) {
	args := m.Called(role)
	return args.Get(0).(model.Role), args.Error(1)
}

func (m *MockRoleRepo) AssignPermissionToRole(roleID string, permissionID string) error {
	args := m.Called(roleID, permissionID)
	return args.Error(0)
}

func TestCreateRole(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockRoleRepo)

	// Define the test input and expected output
	testRole := model.Role{
		ID:   1,
		Name: "Admin",
	}

	// Set up expectations: mock the CreateRole method
	mockRepo.On("CreateRole", testRole).Return(testRole, nil)

	// Create the UseCase with the mocked repository
	useCase := NewRoleUseCase(mockRepo)

	// Call the method under test
	result, err := useCase.CreateRole(testRole)

	// Assert the expectations
	assert.NoError(t, err)
	assert.Equal(t, testRole, result)

	// Assert that the CreateRole method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestCreateRole_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockRoleRepo)

	// Define the test input
	testRole := model.Role{
		ID:   1,
		Name: "Admin",
	}

	// Set up expectations: simulate an error returned by CreateRole
	mockRepo.On("CreateRole", testRole).Return(model.Role{}, errors.New("failed to create role"))

	// Create the UseCase with the mocked repository
	useCase := NewRoleUseCase(mockRepo)

	// Call the method under test
	result, err := useCase.CreateRole(testRole)

	// Assert that an error occurred
	assert.Error(t, err)
	assert.Equal(t, model.Role{}, result)

	// Assert that the CreateRole method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestAssignPermissionToRole(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockRoleRepo)

	// Define the test input
	roleID := "1"
	permissionID := "101"

	// Set up expectations: mock the AssignPermissionToRole method
	mockRepo.On("AssignPermissionToRole", roleID, permissionID).Return(nil)

	// Create the UseCase with the mocked repository
	useCase := NewRoleUseCase(mockRepo)

	// Call the method under test
	err := useCase.AssignPermissionToRole(roleID, permissionID)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert that the AssignPermissionToRole method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestAssignPermissionToRole_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockRoleRepo)

	// Define the test input
	roleID := "1"
	permissionID := "101"

	// Set up expectations: simulate an error returned by AssignPermissionToRole
	mockRepo.On("AssignPermissionToRole", roleID, permissionID).Return(errors.New("failed to assign permission"))

	// Create the UseCase with the mocked repository
	useCase := NewRoleUseCase(mockRepo)

	// Call the method under test
	err := useCase.AssignPermissionToRole(roleID, permissionID)

	// Assert that an error occurred
	assert.Error(t, err)

	// Assert that the AssignPermissionToRole method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}
