package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionStructFields(t *testing.T) {
	// Check if the Permission struct has the correct fields and types
	permissionType := reflect.TypeOf(Permission{})

	// Check the ID field
	idField, idFound := permissionType.FieldByName("ID")
	assert.True(t, idFound, "ID field should be present")
	assert.Equal(t, "uint", idField.Type.Name(), "ID field should be of type uint")
	assert.Contains(t, idField.Tag.Get("gorm"), "primaryKey", "ID field should have primaryKey tag")

	// Check the Name field
	nameField, nameFound := permissionType.FieldByName("Name")
	assert.True(t, nameFound, "Name field should be present")
	assert.Equal(t, "string", nameField.Type.Name(), "Name field should be of type string")
	assert.Contains(t, nameField.Tag.Get("gorm"), "uniqueIndex", "Name field should have uniqueIndex tag")
	assert.Contains(t, nameField.Tag.Get("gorm"), "type:varchar(100)", "Name field should have varchar(100) tag")
	assert.Equal(t, "name", nameField.Tag.Get("json"), "Name field should have json tag 'name'")
}

func TestPermissionJSONMarshaling(t *testing.T) {
	// Test JSON marshaling for the Permission struct
	perm := Permission{ID: 1, Name: "manage_users"}

	expectedJSON := `{"ID":1,"name":"manage_users"}`
	actualJSON, err := json.Marshal(perm)
	assert.NoError(t, err, "JSON marshaling should not produce an error")
	assert.JSONEq(t, expectedJSON, string(actualJSON), "JSON output does not match expected format")
}

func TestPermissionDefaultValues(t *testing.T) {
	// Test that default values of a new Permission struct are as expected
	perm := Permission{}

	assert.Equal(t, uint(0), perm.ID, "Default ID should be 0")
	assert.Equal(t, "", perm.Name, "Default Name should be an empty string")
}
