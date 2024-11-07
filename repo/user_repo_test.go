package repo

import (
	"errors"
	"fmt"
	"go-multirole/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (d *UserRepositoryMock) CreateUser(user model.User) (model.User, error) {
	args := d.Mock.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (d *UserRepositoryMock) LoginUser(inputUser model.User) (model.User, error) {
	args := d.Mock.Called(inputUser)
	return args.Get(0).(model.User), args.Error(1)
}

func (d *UserRepositoryMock) CheckUserPermission(userID string, permissionName string) (bool, error) {
	args := d.Mock.Called(userID, permissionName)
	return args.Bool(0), args.Error(1)
}

// AssignRoleToUser mocks the AssignRoleToUser method of the UserRepository
func (d *UserRepositoryMock) AssignRoleToUser(userId string, roleID string) error {
	args := d.Mock.Called(userId, roleID)
	return args.Error(0)
}

func TestCreateUser_Success(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	testUser := model.User{
		Username: "testuser",
		Password: "password123", // The password will be hashed internally in CreateUser
	}

	// Mock the behavior: return the test user and no error on CreateUser
	repoMock.Mock.On("CreateUser", testUser).Return(testUser, nil)

	// Act
	createdUser, err := repoMock.CreateUser(testUser)

	// Assert
	assert.NoError(t, err)                                   // No error should occur
	assert.Equal(t, testUser.Username, createdUser.Username) // Usernames should match
	repoMock.Mock.AssertExpectations(t)                      // Check all expectations were met
}

func TestCreateUser_Failure(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	testUser := model.User{
		Username: "testuser",
		Password: "password123", // The password will be hashed internally in CreateUser
	}

	// Mock the behavior: simulate a database error on CreateUser
	repoMock.Mock.On("CreateUser", testUser).Return(testUser, errors.New("database error"))

	// Act
	createdUser, err := repoMock.CreateUser(testUser)

	// Assert
	assert.Error(t, err)                                     // Error should occur
	assert.EqualError(t, err, "database error")              // Error should match expected message
	assert.Equal(t, testUser.Username, createdUser.Username) // Usernames should match
	repoMock.Mock.AssertExpectations(t)                      // Check all expectations were met
}

func TestLoginUser_Success(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	testUser := model.User{
		Username: "testuser",
		Password: "hashedPassword123", // Assume the password is already hashed
	}

	// Simulate the behavior: return the test user and no error when the username is found
	repoMock.Mock.On("LoginUser", testUser).Return(testUser, nil)

	// Act
	returnedUser, err := repoMock.LoginUser(testUser)

	// Assert
	assert.NoError(t, err)                                    // No error should occur
	assert.Equal(t, testUser.Username, returnedUser.Username) // Usernames should match
	repoMock.Mock.AssertExpectations(t)                       // Check all expectations were met
}

func TestLoginUser_UserNotFound(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	testUser := model.User{
		Username: "testuser",
		Password: "hashedPassword123", // Assume the password is already hashed
	}

	// Simulate the behavior: user is not found (gorm.ErrRecordNotFound) and return an error
	repoMock.Mock.On("LoginUser", testUser).Return(model.User{}, errors.New("user not found"))

	// Act
	returnedUser, err := repoMock.LoginUser(testUser)

	// Assert
	assert.Error(t, err)                        // Error should occur
	assert.EqualError(t, err, "user not found") // Error message should be "user not found"
	assert.Equal(t, model.User{}, returnedUser) // Returned user should be empty
	repoMock.Mock.AssertExpectations(t)         // Check all expectations were met
}

func TestLoginUser_DatabaseError(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	testUser := model.User{
		Username: "testuser",
		Password: "hashedPassword123", // Assume the password is already hashed
	}

	// Simulate a database error when querying the user
	repoMock.Mock.On("LoginUser", testUser).Return(model.User{}, fmt.Errorf("database error"))

	// Act
	returnedUser, err := repoMock.LoginUser(testUser)

	// Assert
	assert.Error(t, err)                        // Error should occur
	assert.EqualError(t, err, "database error") // Error message should be "database error"
	assert.Equal(t, model.User{}, returnedUser) // Returned user should be empty
	repoMock.Mock.AssertExpectations(t)         // Check all expectations were met
}

func TestAssignRoleToUser_Success(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)

	// Simulate behavior: the user and role are found and the role is assigned successfully
	repoMock.Mock.On("AssignRoleToUser", "1", "1").Return(nil) // Expect the method to succeed

	// Act
	err := repoMock.AssignRoleToUser("1", "1")

	// Assert
	assert.NoError(t, err)              // No error should occur
	repoMock.Mock.AssertExpectations(t) // Check that all expectations were met
}

func TestAssignRoleToUser_UserNotFound(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)

	// Simulate behavior: user not found in the database
	repoMock.Mock.On("AssignRoleToUser", "1", "1").Return(errors.New("user not found"))

	// Act
	err := repoMock.AssignRoleToUser("1", "1")

	// Assert
	assert.Error(t, err)                        // Error should occur
	assert.EqualError(t, err, "user not found") // Error message should be "user not found"
	repoMock.Mock.AssertExpectations(t)         // Check that all expectations were met
}

func TestAssignRoleToUser_RoleNotFound(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)

	// Simulate behavior: role not found in the database
	repoMock.Mock.On("AssignRoleToUser", "1", "1").Return(errors.New("role not found"))

	// Act
	err := repoMock.AssignRoleToUser("1", "1")

	// Assert
	assert.Error(t, err)                        // Error should occur
	assert.EqualError(t, err, "role not found") // Error message should be "role not found"
	repoMock.Mock.AssertExpectations(t)         // Check that all expectations were met
}

func TestAssignRoleToUser_DatabaseError(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)

	// Simulate behavior: database error while assigning role
	repoMock.Mock.On("AssignRoleToUser", "1", "1").Return(errors.New("database error"))

	// Act
	err := repoMock.AssignRoleToUser("1", "1")

	// Assert
	assert.Error(t, err)                        // Error should occur
	assert.EqualError(t, err, "database error") // Error message should be "database error"
	repoMock.Mock.AssertExpectations(t)         // Check that all expectations were met
}

func TestCheckUserPermission_Success(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	permissionName := "read_permission"

	// Simulate behavior: the user exists and the permission is found
	repoMock.Mock.On("CheckUserPermission", "1", permissionName).Return(true, nil)

	// Act
	hasPermission, err := repoMock.CheckUserPermission("1", permissionName)

	// Assert
	assert.NoError(t, err)              // No error should occur
	assert.True(t, hasPermission)       // The user should have the permission
	repoMock.Mock.AssertExpectations(t) // Check that all expectations were met
}

func TestCheckUserPermission_PermissionNotFound(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)

	permissionName := "read_permission"

	// Simulate behavior: the user exists but the permission is not found
	repoMock.Mock.On("CheckUserPermission", "1", permissionName).Return(false, nil)

	// Act
	hasPermission, err := repoMock.CheckUserPermission("1", permissionName)

	// Assert
	assert.NoError(t, err)              // No error should occur
	assert.False(t, hasPermission)      // The user should not have the permission
	repoMock.Mock.AssertExpectations(t) // Check that all expectations were met
}

func TestCheckUserPermission_UserNotFound(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	permissionName := "read_permission"

	// Simulate behavior: user not found in the database
	repoMock.Mock.On("CheckUserPermission", "1", permissionName).Return(false, errors.New("user not found"))

	// Act
	hasPermission, err := repoMock.CheckUserPermission("1", permissionName)

	// Assert
	assert.Error(t, err)                        // Error should occur
	assert.EqualError(t, err, "user not found") // Error message should be "user not found"
	assert.False(t, hasPermission)              // The user should not have the permission
	repoMock.Mock.AssertExpectations(t)         // Check that all expectations were met
}

func TestCheckUserPermission_DatabaseError(t *testing.T) {
	// Arrange
	repoMock := new(UserRepositoryMock)
	permissionName := "read_permission"

	// Simulate behavior: database error while checking permission
	repoMock.Mock.On("CheckUserPermission", "1", permissionName).Return(false, errors.New("database error"))

	// Act
	hasPermission, err := repoMock.CheckUserPermission("1", permissionName)

	// Assert
	assert.Error(t, err)                        // Error should occur
	assert.EqualError(t, err, "database error") // Error message should be "database error"
	assert.False(t, hasPermission)              // The user should not have the permission
	repoMock.Mock.AssertExpectations(t)         // Check that all expectations were met
}
