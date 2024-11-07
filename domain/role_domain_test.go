package domain

import (
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for RoleRepo interface
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

// Mock for RoleUseCase interface
type MockRoleUseCase struct {
	mock.Mock
}

func (m *MockRoleUseCase) CreateRole(role model.Role) (model.Role, error) {
	args := m.Called(role)
	return args.Get(0).(model.Role), args.Error(1)
}

func (m *MockRoleUseCase) AssignPermissionToRole(roleID string, permissionID string) error {
	args := m.Called(roleID, permissionID)
	return args.Error(0)
}

// Unit Test for RoleRepo interface
func TestRoleRepo(t *testing.T) {
	mockRepo := new(MockRoleRepo)

	// Test: Create Role
	t.Run("Create Role", func(t *testing.T) {
		role := model.Role{ID: 1, Name: "admin"}
		mockRepo.On("CreateRole", role).Return(role, nil)

		createdRole, err := mockRepo.CreateRole(role)

		assert.NoError(t, err)
		assert.Equal(t, role, createdRole)
		mockRepo.AssertExpectations(t)
	})

	// Test: Assign Permission to Role
	t.Run("Assign Permission to Role", func(t *testing.T) {
		roleID := "1"
		permissionID := "read_permission"
		mockRepo.On("AssignPermissionToRole", roleID, permissionID).Return(nil)

		err := mockRepo.AssignPermissionToRole(roleID, permissionID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// Unit Test for RoleUseCase interface
func TestRoleUseCase(t *testing.T) {
	mockUseCase := new(MockRoleUseCase)

	// Test: Create Role
	t.Run("Create Role", func(t *testing.T) {
		role := model.Role{ID: 1, Name: "admin"}
		mockUseCase.On("CreateRole", role).Return(role, nil)

		createdRole, err := mockUseCase.CreateRole(role)

		assert.NoError(t, err)
		assert.Equal(t, role, createdRole)
		mockUseCase.AssertExpectations(t)
	})

	// Test: Assign Permission to Role
	t.Run("Assign Permission to Role", func(t *testing.T) {
		roleID := "1"
		permissionID := "read_permission"
		mockUseCase.On("AssignPermissionToRole", roleID, permissionID).Return(nil)

		err := mockUseCase.AssignPermissionToRole(roleID, permissionID)

		assert.NoError(t, err)
		mockUseCase.AssertExpectations(t)
	})
}
