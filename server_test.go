package main

import (
	"github.com/PolkaSpots/puffin-api-v2/db"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	// "net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {

	homeHandle := HomeHandler()
	test := GenerateHandleTester(t, homeHandle)

	w := test("GET", "")

	log.Print(w.Body)

	if w.Code != http.StatusOK {
		t.Errorf(
			"GET / returned %v. Expected %v",
			w.Code,
			http.StatusOK,
		)
	}

	expected := "Gorilla!\n"
	actual, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(actual) {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}

func TestLoginHandler(t *testing.T) {

	mockDb := db.NewMockDbManager(false)

	// mockDb.RegisterUser("test user", hash)

	registerHandle := RegHandler(mockDb)
	test := GenerateHandleTester(t, registerHandle)

	mockDb.RegisterUser("test userss")

	goodParams := `{
		"username": "test user",
		"password1": "test pass",
		"password2": "test pass"
	}`

	w := httptest.NewRecorder()

	w = test("POST", goodParams)
	if w.Code != http.StatusCreated {
		t.Errorf(
			"POST /login: good input returned %v. Expected %v.",
			w.Code,
			http.StatusCreated,
		)
	}
}
