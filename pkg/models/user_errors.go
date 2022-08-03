package models

import (
	"fmt"
)

type Entity interface{}

type UserExistsError struct {
	Email string
}

func (err *UserExistsError) Error() string {
	return fmt.Sprintf("User with email %s already exists", err.Email)
}

type UserNotFoundError struct {
	Entity Entity
}

func (err *UserNotFoundError) Error() string {
	return fmt.Sprintf("User with entity %s not found", err.Entity)
}

type WrongPasswordError struct {
	Email string
}

func (err *WrongPasswordError) Error() string {
	return fmt.Sprintf("Wrong password for user with email %s", err.Email)
}
