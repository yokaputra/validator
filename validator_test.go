package validator

import (
	"testing"
)

func TestValidatorImpl_Validate(t *testing.T) {
	type testStruct struct {
		Name   string `json:"name" validate:"required"`
		Email  string `json:"email" validate:"required,email"`
		Age    int    `json:"age" validate:"required,gt=0"`
		Gender string `json:"gender" validate:"required,oneof='MAN' 'WOMEN'"`
	}

	tests := []struct {
		name           string
		input          interface{}
		expectedErrMsg string
	}{
		{
			name: "valid input",
			input: testStruct{
				Name:   "John Doe",
				Email:  "johndoe@example.com",
				Age:    30,
				Gender: "MAN",
			},
			expectedErrMsg: "",
		},
		{
			name: "missing required field",
			input: testStruct{
				Email:  "johndoe@example.com",
				Age:    30,
				Gender: "MAN",
			},
			expectedErrMsg: "validation failed on field name: required",
		},
		{
			name: "invalid email format",
			input: testStruct{
				Name:   "John Doe",
				Email:  "johndoe",
				Age:    30,
				Gender: "MAN",
			},
			expectedErrMsg: "validation failed on field email: email",
		},
		{
			name: "invalid age",
			input: testStruct{
				Name:   "John Doe",
				Email:  "johndoe@example.com",
				Age:    -1,
				Gender: "MAN",
			},
			expectedErrMsg: "validation failed on field age: gt",
		},
		{
			name: "invalid gender",
			input: testStruct{
				Name:   "John Doe",
				Email:  "johndoe@example.com",
				Age:    30,
				Gender: "OTHER",
			},
			expectedErrMsg: "validation failed on field gender: oneof, Expected: 'MAN' 'WOMEN', Actual: OTHER",
		},
	}

	validator := NewValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.input)
			if err != nil && err.Error() != tt.expectedErrMsg {
				t.Errorf("Expected error message: %s, but got: %s", tt.expectedErrMsg, err.Error())
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Errorf("Expected error message: %s, but got nil", tt.expectedErrMsg)
			}
		})
	}
}
