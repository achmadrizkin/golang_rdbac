package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleStructFields(t *testing.T) {
	// Check if the Role struct has the correct fields and types
	roleType := reflect.TypeOf(Role{})

	// Check the ID field
	idField, idFound := roleType.FieldByName("ID")
	assert.True(t, idFound, "ID field should be present")
	assert.Equal(t, "uint", idField.Type.Name(), "ID field should be of type uint")
	assert.Contains(t, idField.Tag.Get("gorm"), "primaryKey", "ID field should have primaryKey tag")

	// Check the Name field
	nameField, nameFound := roleType.FieldByName("Name")
	assert.True(t, nameFound, "Name field should be present")
	assert.Equal(t, "string", nameField.Type.Name(), "Name field should be of type string")
	assert.Contains(t, nameField.Tag.Get("gorm"), "uniqueIndex", "Name field should have uniqueIndex tag")
	assert.Contains(t, nameField.Tag.Get("gorm"), "type:varchar(100)", "Name field should have varchar(100) tag")
	assert.Equal(t, "name", nameField.Tag.Get("json"), "Name field should have json tag 'name'")

	// Check the Permissions field
	permissionsField, permissionsFound := roleType.FieldByName("Permissions")
	assert.True(t, permissionsFound, "Permissions field should be present")
	assert.Equal(t, "slice", permissionsField.Type.Kind().String(), "Permissions field should be a slice type")
	assert.Contains(t, permissionsField.Tag.Get("gorm"), "many2many:role_permissions", "Permissions field should have many2many relationship tag")
	assert.Equal(t, "permissions", permissionsField.Tag.Get("json"), "Permissions field should have json tag 'permissions'")
}

func TestRoleJSONMarshaling(t *testing.T) {
	// Test JSON marshaling with permissions populated
	role := Role{
		ID:   1,
		Name: "Admin",
		Permissions: []Permission{
			{ID: 1, Name: "read"},
			{ID: 2, Name: "write"},
		},
	}

	expectedJSON := `{
		"ID": 1,
		"name": "Admin",
		"permissions": [
			{"ID": 1, "name": "read"},
			{"ID": 2, "name": "write"}
		]
	}`
	actualJSON, err := json.Marshal(role)
	assert.NoError(t, err, "JSON marshaling should not produce an error")
	assert.JSONEq(t, expectedJSON, string(actualJSON), "JSON output with permissions does not match expected format")
}

func TestRoleDefaultValues(t *testing.T) {
	// Test that default values of a new Role struct are as expected
	role := Role{}

	assert.Equal(t, uint(0), role.ID, "Default ID should be 0")
	assert.Equal(t, "", role.Name, "Default Name should be an empty string")
	assert.Nil(t, role.Permissions, "Default Permissions should be nil")
}
