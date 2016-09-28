package db

import (
	"errors"
)

type MockDbManager struct {
	registrationError bool
	users             map[string]string
	// files             map[string]*common.CartonFile
}

func NewMockDbManager(registrationError bool) *MockDbManager {
	db := MockDbManager{}
	db.registrationError = registrationError
	db.users = make(map[string]string)
	// db.files = make(map[string]*common.CartonFile)
	return &db
}

// func (db *MockDbManager) Close() error {
// 	return nil
// }

func (db *MockDbManager) RegisterUser(user string) error {
	if db.registrationError {
		return errors.New("Registration error")
	}
	db.users[user] = "cheese"
	return nil
}
