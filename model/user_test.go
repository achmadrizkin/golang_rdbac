package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserStructFields(t *testing.T) {
	// Check if the User struct has the correct fields and types
	userType := reflect.TypeOf(User{})

	// Check the ID field
	idField, idFound := userType.FieldByName("ID")
	assert.True(t, idFound, "ID field should be present")
	assert.Equal(t, "uint", idField.Type.Name(), "ID field should be of type uint")
	assert.Contains(t, idField.Tag.Get("gorm"), "primaryKey", "ID field should have primaryKey tag")

	// Check the Username field
	usernameField, usernameFound := userType.FieldByName("Username")
	assert.True(t, usernameFound, "Username field should be present")
	assert.Equal(t, "string", usernameField.Type.Name(), "Username field should be of type string")
	assert.Contains(t, usernameField.Tag.Get("gorm"), "uniqueIndex", "Username field should have uniqueIndex tag")
	assert.Contains(t, usernameField.Tag.Get("gorm"), "type:varchar(100)", "Username field should have varchar(100) tag")
	assert.Equal(t, "username", usernameField.Tag.Get("json"), "Username field should have json tag 'username'")

	// Check the Password field
	passwordField, passwordFound := userType.FieldByName("Password")
	assert.True(t, passwordFound, "Password field should be present")
	assert.Equal(t, "string", passwordField.Type.Name(), "Password field should be of type string")
	assert.Contains(t, passwordField.Tag.Get("gorm"), "type:varchar(100)", "Password field should have varchar(100) tag")
	assert.Equal(t, "password", passwordField.Tag.Get("json"), "Password field should have json tag 'password'")

	// Check the Roles field
	rolesField, rolesFound := userType.FieldByName("Roles")
	assert.True(t, rolesFound, "Roles field should be present")
	assert.Equal(t, "slice", rolesField.Type.Kind().String(), "Roles field should be a slice type")
	assert.Contains(t, rolesField.Tag.Get("gorm"), "many2many:user_roles", "Roles field should have many2many relationship tag")
	assert.Equal(t, "roles", rolesField.Tag.Get("json"), "Roles field should have json tag 'roles'")
}

func TestUserJSONMarshaling(t *testing.T) {
	// Test JSON marshaling with roles populated
	user := User{
		ID:       1,
		Username: "johndoe",
		Password: "password123",
		Roles: []Role{
			{ID: 1, Name: "Admin"},
			{ID: 2, Name: "User"},
		},
	}

	expectedJSON := `{
		"ID": 1,
		"username": "johndoe",
		"password": "password123",
		"roles": [
			{"ID": 1, "name": "Admin", "permissions": null},
			{"ID": 2, "name": "User", "permissions": null}
		]
	}`
	actualJSON, err := json.Marshal(user)
	assert.NoError(t, err, "JSON marshaling should not produce an error")
	assert.JSONEq(t, expectedJSON, string(actualJSON), "JSON output with roles does not match expected format")
}

func TestUserDefaultValues(t *testing.T) {
	// Test that default values of a new User struct are as expected
	user := User{}

	assert.Equal(t, uint(0), user.ID, "Default ID should be 0")
	assert.Equal(t, "", user.Username, "Default Username should be an empty string")
	assert.Equal(t, "", user.Password, "Default Password should be an empty string")
	assert.Nil(t, user.Roles, "Default Roles should be nil")
}
