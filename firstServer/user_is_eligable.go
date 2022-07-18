package main

import (
	"errors"
	"fmt"
)

func userIsEligable(email, password string, age int) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	if password == "" {
		return errors.New("password cannot be empty")
	}

	const minAge = 18

	if age < 18 {
		return fmt.Errorf("age must be atleast %v years old", minAge)
	}

	return nil
}