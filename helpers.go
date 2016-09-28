package main

import (
	"net/http"
	"strings"

	"net/http/httptest"
	"testing"
)

type HandleTester func(
	method string,
	params string,
) *httptest.ResponseRecorder

func GenerateHandleTester(
	t *testing.T,
	handleFunc http.Handler,
) HandleTester {

	// Given a method type ("GET", "POST", etc) and
	// parameters, serve the response against the handler
	// and return the ResponseRecorder.

	return func(
		method string,
		params string,
	) *httptest.ResponseRecorder {

		req, err := http.NewRequest(
			method,
			"",
			strings.NewReader(params),
		)
		if err != nil {
			t.Errorf("%v", err)
		}
		req.Header.Set(
			"Content-Type",
			"application/json",
		)
		req.Body.Close()
		w := httptest.NewRecorder()
		handleFunc.ServeHTTP(w, req)
		return w
	}
}
