package db

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	// tempDb := getTempFile()
	// if tempDb == "" {
	// 	t.Skip("Cannot create temp file")
	// }

	m, _ := NewDatastore()

	// if m.IsUser("fake user") {
	// 	t.Errorf("IsUsers returns true for 'fake user'")
	// }
	// if m.GetPwdHash("fake user") != nil {
	// 	t.Errorf("GetPwdHash returns non nil value for 'fake user'")
	// }

	err := m.RegisterUser("test user")
	if err != nil {
		t.Errorf("failed to register 'test user'")
	}
}
