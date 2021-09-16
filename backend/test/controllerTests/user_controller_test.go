package controllertests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amr9mohamed/mainApp/api/controllers"
	"github.com/amr9mohamed/mainApp/api/models"
	"github.com/go-playground/assert/v2"
)

var router = controllers.Server{}.Router

func TestGetUsers(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	router.ServeHTTP(w, req)
	var users []models.User
	err = json.Unmarshal((w.Body.Bytes()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(users), 1007616)
}

func TestGetUsersByCountry(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/:country/:pageNumber/:pageSize/", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	router.ServeHTTP(w, req)

	var users []models.User
	err = json.Unmarshal((w.Body.Bytes()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, len(users), 0)
}

func TestGetDistinctCountries(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/countries", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	router.ServeHTTP(w, req)

	var users []models.User
	err = json.Unmarshal((w.Body.Bytes()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(users), 5)
}

func TestGetCountryFrequency(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/frequency", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	router.ServeHTTP(w, req)

	results := []models.CountryFrequency{}
	err = json.Unmarshal((w.Body.Bytes()), &results)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(results), 5)
}
