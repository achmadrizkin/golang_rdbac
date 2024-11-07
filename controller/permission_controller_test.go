package controller

import (
	"bytes"
	"go-multirole/model"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPermissionUseCase is a mock implementation of the PermissionUseCase interface
type MockPermissionUseCase struct {
	mock.Mock
}

func (m *MockPermissionUseCase) CreatePermission(permission model.Permission) (model.Permission, error) {
	args := m.Called(permission)
	return args.Get(0).(model.Permission), args.Error(1)
}

// Unit tests for PermissionController
func TestPermissionController(t *testing.T) {
	mockUseCase := new(MockPermissionUseCase)
	permissionController := NewPermissionController(mockUseCase)

	t.Run("Create permission successfully", func(t *testing.T) {
		// Define a mock permission with ID 0 since it will be generated later
		mockPermission := model.Permission{ID: 0, Name: "admin_access"}

		// Mock the return of CreatePermission with the mock permission and no error
		mockUseCase.On("CreatePermission", mockPermission).Return(model.Permission{ID: 1, Name: "admin_access"}, nil)

		// Create a test HTTP request with a JSON body that matches the permission model
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/permission", bytes.NewBufferString(`{"name":"admin_access"}`))

		// Call the CreatePermission function
		permissionController.CreatePermission(c)

		// Assert the response status is Created and the message is as expected
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Created permission success")

		// Verify mock expectations
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Create permission with invalid JSON", func(t *testing.T) {
		// Create a test HTTP request with invalid JSON
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/permission", bytes.NewBufferString(`{"name":`)) // invalid JSON

		// Call the CreatePermission function
		permissionController.CreatePermission(c)

		// Assert the response status is BadRequest
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "unexpected EOF")
	})
}
