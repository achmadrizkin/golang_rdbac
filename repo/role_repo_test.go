package repo

import (
	"errors"
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RoleRepositoryMock struct {
	Mock mock.Mock
}

func (repository *RoleRepositoryMock) CreateRole(role model.Role) (model.Role, error) {
	args := repository.Mock.Called(role)
	// Return the arguments as model.Role and error.
	return args.Get(0).(model.Role), args.Error(1)
}

func (repository *RoleRepositoryMock) AssignPermissionToRole(roleID string, permissionID string) error {
	args := repository.Mock.Called(roleID, permissionID)
	return args.Error(0)
}

func TestCreateRole_Success(t *testing.T) {
	// Arrange
	repoMock := new(RoleRepositoryMock)
	testRole := model.Role{Name: "Admin"}

	// Mock the behavior: return the test role and no error on CreateRole
	repoMock.Mock.On("CreateRole", testRole).Return(testRole, nil)

	// Act
	createdRole, err := repoMock.CreateRole(testRole)

	// Assert
	assert.NoError(t, err)                           // No error should occur
	assert.Equal(t, testRole.Name, createdRole.Name) // Role name should match
	repoMock.Mock.AssertExpectations(t)              // Check all expectations were met
}

func TestCreateRole_Failure(t *testing.T) {
	// Arrange
	repoMock := new(RoleRepositoryMock)
	testRole := model.Role{Name: "Admin"}
	expectedError := errors.New("database error")

	// Mock the behavior: return an error on CreateRole
	repoMock.Mock.On("CreateRole", testRole).Return(model.Role{}, expectedError)

	// Act
	createdRole, err := repoMock.CreateRole(testRole)

	// Assert
	assert.Error(t, err)                        // An error should occur
	assert.EqualError(t, err, "database error") // Error message should match
	assert.Equal(t, model.Role{}, createdRole)  // Created role should be empty
	repoMock.Mock.AssertExpectations(t)         // Check all expectations were met
}

func TestAssignPermissionToRole_Success(t *testing.T) {
	// Arrange
	repoMock := new(RoleRepositoryMock)
	roleID := "1"
	permissionID := "10"

	// Mock the behavior: no error returned on AssignPermissionToRole
	repoMock.Mock.On("AssignPermissionToRole", roleID, permissionID).Return(nil)

	// Act
	err := repoMock.AssignPermissionToRole(roleID, permissionID)

	// Assert
	assert.NoError(t, err)              // Expect no error
	repoMock.Mock.AssertExpectations(t) // Check all expectations were met
}

func TestAssignPermissionToRole_Failure_RoleNotFound(t *testing.T) {
	// Arrange
	repoMock := new(RoleRepositoryMock)
	roleID := "1"
	permissionID := "10"
	expectedError := errors.New("role not found")

	// Mock the behavior: return "role not found" error
	repoMock.Mock.On("AssignPermissionToRole", roleID, permissionID).Return(expectedError)

	// Act
	err := repoMock.AssignPermissionToRole(roleID, permissionID)

	// Assert
	assert.Error(t, err)                        // Expect an error
	assert.EqualError(t, err, "role not found") // Error message should match
	repoMock.Mock.AssertExpectations(t)         // Check all expectations were met
}

func TestAssignPermissionToRole_Failure_PermissionNotFound(t *testing.T) {
	// Arrange
	repoMock := new(RoleRepositoryMock)
	roleID := "1"
	permissionID := "10"
	expectedError := errors.New("permission not found")

	// Mock the behavior: return "permission not found" error
	repoMock.Mock.On("AssignPermissionToRole", roleID, permissionID).Return(expectedError)

	// Act
	err := repoMock.AssignPermissionToRole(roleID, permissionID)

	// Assert
	assert.Error(t, err)                              // Expect an error
	assert.EqualError(t, err, "permission not found") // Error message should match
	repoMock.Mock.AssertExpectations(t)               // Check all expectations were met
}

func TestAssignPermissionToRole_Failure_AssociationError(t *testing.T) {
	// Arrange
	repoMock := new(RoleRepositoryMock)
	roleID := "1"
	permissionID := "10"
	expectedError := errors.New("failed to associate permission with role")

	// Mock the behavior: return "association error" when appending the permission
	repoMock.Mock.On("AssignPermissionToRole", roleID, permissionID).Return(expectedError)

	// Act
	err := repoMock.AssignPermissionToRole(roleID, permissionID)

	// Assert
	assert.Error(t, err)                                                  // Expect an error
	assert.EqualError(t, err, "failed to associate permission with role") // Error message should match
	repoMock.Mock.AssertExpectations(t)                                   // Check all expectations were met
}
