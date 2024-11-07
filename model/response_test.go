package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseStructFields(t *testing.T) {
	// Check if the Response struct has the correct fields and types
	responseType := reflect.TypeOf(Response{})

	// Check the StatusCode field
	statusCodeField, statusCodeFound := responseType.FieldByName("StatusCode")
	assert.True(t, statusCodeFound, "StatusCode field should be present")
	assert.Equal(t, "int", statusCodeField.Type.Name(), "StatusCode field should be of type int")
	assert.Equal(t, "status_code", statusCodeField.Tag.Get("json"), "StatusCode field should have json tag 'status_code'")

	// Check the Message field
	messageField, messageFound := responseType.FieldByName("Message")
	assert.True(t, messageFound, "Message field should be present")
	assert.Equal(t, "string", messageField.Type.Name(), "Message field should be of type string")
	assert.Equal(t, "message", messageField.Tag.Get("json"), "Message field should have json tag 'message'")

	// Check the Data field
	dataField, dataFound := responseType.FieldByName("Data")
	assert.True(t, dataFound, "Data field should be present")
	assert.Equal(t, "interface {}", dataField.Type.String(), "Data field should be of type interface{}")
	assert.Contains(t, dataField.Tag.Get("json"), "data", "Data field should have json tag 'data'")
	assert.Contains(t, dataField.Tag.Get("json"), "omitempty", "Data field should have omitempty tag")
}

func TestResponseJSONMarshalingWithAndWithoutData(t *testing.T) {
	// Test JSON marshaling with Data field set
	respWithData := Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       map[string]string{"key": "value"},
	}

	expectedJSONWithData := `{"status_code":200,"message":"Success","data":{"key":"value"}}`
	actualJSONWithData, err := json.Marshal(respWithData)
	assert.NoError(t, err, "JSON marshaling with Data should not produce an error")
	assert.JSONEq(t, expectedJSONWithData, string(actualJSONWithData), "JSON output with Data does not match expected format")

	// Test JSON marshaling without Data field (should omit data)
	respWithoutData := Response{
		StatusCode: 404,
		Message:    "Not Found",
	}

	expectedJSONWithoutData := `{"status_code":404,"message":"Not Found"}`
	actualJSONWithoutData, err := json.Marshal(respWithoutData)
	assert.NoError(t, err, "JSON marshaling without Data should not produce an error")
	assert.JSONEq(t, expectedJSONWithoutData, string(actualJSONWithoutData), "JSON output without Data does not match expected format")
}

func TestResponseDefaultValues(t *testing.T) {
	// Test that default values of a new Response struct are as expected
	resp := Response{}

	assert.Equal(t, 0, resp.StatusCode, "Default StatusCode should be 0")
	assert.Equal(t, "", resp.Message, "Default Message should be an empty string")
	assert.Nil(t, resp.Data, "Default Data should be nil")
}
