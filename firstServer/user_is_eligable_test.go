package main

import (
	"testing"
	"errors"
)

func TestUserIsEligible(t *testing.T) {

	var tests = []struct {
		email       string
		password    string
		age         int
		expectedErr error
	}{
		{
			email:       "test@example.com",
			password:    "12345",
			age:         18,
			expectedErr:  nil,
		},
		{
			email:       "test@example.com",
			password:    "12345",
			age:         18,
			expectedErr: nil,
		},
		{
			email:       "test@example.com",
			password:    "12345",
			age:         18,
			expectedErr: nil,
		},
		{
			email:       "test@example.com",
			password:    "12345",
			age:         19,
			expectedErr: errors.New("age must be at least 18 years old"),
		},
	}

	for _, tt := range tests {
		err := userIsEligable(tt.email, tt.password, tt.age)
		errString := ""
		expectedErrString := ""
		if err != nil {
			errString = err.Error()
		}
		if tt.expectedErr != nil {
			expectedErrString = tt.expectedErr.Error()
		}
		if errString != expectedErrString {
			t.Errorf("got %s, want %s", errString, expectedErrString)
		}
	}


}

