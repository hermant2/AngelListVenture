package test

import (
	"encoding/json"
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

func AssertErrorResponse(assertions *assert.Assertions, recorder *httptest.ResponseRecorder, expectedError apperror.Standard) {
	var body interface{}
	json.NewDecoder(recorder.Result().Body).Decode(&body)
	errorMap := body.(map[string]interface{})["error"].(map[string]interface{})
	assertions.Equal(int(errorMap["status"].(float64)), expectedError.Status)
	assertions.Equal(int(errorMap["code"].(float64)), int(expectedError.Code))
	assertions.Equal(errorMap["description"], expectedError.Description)
}
