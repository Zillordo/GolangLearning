package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}
	if err == nil {
		t.Error("we were expecting na error when user id is 0")
	}
	if err != nil && err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when user is not found")
	}
	if err != nil && err.Code != "not_found" {
		t.Error("we were expecting not_found when user is not found")
	}
}

func TestGetUserNoError(t *testing.T) {
	// initialization
	users = map[int64]*User{
		1: {Id: 1, FirstName: "TestFirst", LastName: "TestLast", Email: "test@test.com"},
	}

	// tests
	user, err := GetUser(1)
	if err != nil {
		t.Error("we were not expecting a error if user exists")
	}
	if user == nil {
		t.Error("we were expecting a user")
	}
	if user != nil && user.Id != 1 {
		t.Error("we were expecting a user id to equal 1")
	}
	if user != nil && user.FirstName != "TestFirst" {
		t.Error("we were expecting a FirstName to equal TestFirst")
	}
	if user != nil && user.LastName != "TestLast" {
		t.Error("we were expecting a FirstName to equal TestLast")
	}
	if user != nil && user.Email != "test@test.com" {
		t.Error("we were expecting a FirstName to equal test@test.com")
	}
}
