package domain

import (
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for PermissionRepo interface
type MockPermissionRepo struct {
	mock.Mock
}

func (m *MockPermissionRepo) CreatePermission(permission model.Permission) (model.Permission, error) {
	args := m.Called(permission)
	return args.Get(0).(model.Permission), args.Error(1)
}

// Mock for PermissionUseCase interface
type MockPermissionUseCase struct {
	mock.Mock
}

func (m *MockPermissionUseCase) CreatePermission(permission model.Permission) (model.Permission, error) {
	args := m.Called(permission)
	return args.Get(0).(model.Permission), args.Error(1)
}

// Unit Test for PermissionRepo interface
func TestPermissionRepo(t *testing.T) {
	mockRepo := new(MockPermissionRepo)

	// Test: Create Permission
	t.Run("Create Permission", func(t *testing.T) {
		permission := model.Permission{ID: 1, Name: "read"}
		mockRepo.On("CreatePermission", permission).Return(permission, nil)

		createdPermission, err := mockRepo.CreatePermission(permission)

		assert.NoError(t, err)
		assert.Equal(t, permission, createdPermission)
		mockRepo.AssertExpectations(t)
	})
}

// Unit Test for PermissionUseCase interface
func TestPermissionUseCase(t *testing.T) {
	mockUseCase := new(MockPermissionUseCase)

	// Test: Create Permission
	t.Run("Create Permission", func(t *testing.T) {
		permission := model.Permission{ID: 1, Name: "read"}
		mockUseCase.On("CreatePermission", permission).Return(permission, nil)

		createdPermission, err := mockUseCase.CreatePermission(permission)

		assert.NoError(t, err)
		assert.Equal(t, permission, createdPermission)
		mockUseCase.AssertExpectations(t)
	})
}
