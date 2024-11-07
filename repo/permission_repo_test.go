package repo

import (
	"errors"
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PermissionRepositoryMock struct {
	Mock mock.Mock
}

func (repository *PermissionRepositoryMock) CreatePermission(permission model.Permission) (model.Permission, error) {
	args := repository.Mock.Called(permission)
	// Return the arguments as model.Role and error.
	return args.Get(0).(model.Permission), args.Error(1)
}

func TestCreatePermission_Success(t *testing.T) {
	// Arrange
	repoMock := new(PermissionRepositoryMock)
	testPermission := model.Permission{Name: "read_permission"}

	// Simulate successful permission creation with no error
	repoMock.Mock.On("CreatePermission", testPermission).Return(testPermission, nil)

	// Act
	createdPermission, err := repoMock.CreatePermission(testPermission)

	// Assert
	assert.NoError(t, err)                                       // No error should occur
	assert.Equal(t, testPermission.Name, createdPermission.Name) // Permission name should match
	repoMock.Mock.AssertExpectations(t)                          // Check that all expectations were met
}

func TestCreatePermission_Failure(t *testing.T) {
	// Arrange
	repoMock := new(PermissionRepositoryMock)
	testPermission := model.Permission{Name: "read_permission"}

	// Simulate a database error during permission creation
	repoMock.Mock.On("CreatePermission", testPermission).Return(model.Permission{}, errors.New("database error"))

	// Act
	createdPermission, err := repoMock.CreatePermission(testPermission)

	// Assert
	assert.Error(t, err)                                   // Error should occur
	assert.EqualError(t, err, "database error")            // Error message should be "database error"
	assert.Equal(t, model.Permission{}, createdPermission) // Permission should be empty
	repoMock.Mock.AssertExpectations(t)                    // Check that all expectations were met
}

func TestCreatePermission_EmptyName(t *testing.T) {
	// Arrange
	repoMock := new(PermissionRepositoryMock)
	testPermission := model.Permission{Name: ""} // Empty name

	// Simulate that the method works with empty name
	repoMock.Mock.On("CreatePermission", testPermission).Return(model.Permission{}, errors.New("empty name"))

	// Act
	createdPermission, err := repoMock.CreatePermission(testPermission)

	// Assert
	assert.Error(t, err)                                   // Error should occur
	assert.EqualError(t, err, "empty name")                // Error message should be "empty name"
	assert.Equal(t, model.Permission{}, createdPermission) // Permission should be empty
	repoMock.Mock.AssertExpectations(t)                    // Check that all expectations were met
}
