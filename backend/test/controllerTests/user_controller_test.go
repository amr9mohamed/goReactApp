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

var server = controllers.Server{}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUsers)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal((rr.Body.Bytes()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 1007616)
}

func TestGetUsersByCountry(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/{country}/{pageNumber}/{pageSize}/", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUsersByCountry)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal((rr.Body.Bytes()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.NotEqual(t, len(users), 0)
}

func TestGetDistinctCountries(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/countries", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetDistinctCountries)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal((rr.Body.Bytes()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 5)
}

func TestGetCountryFrequency(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/frequency", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetCountryFrequency)
	handler.ServeHTTP(rr, req)

	results := []models.CountryFrequency{}
	err = json.Unmarshal((rr.Body.Bytes()), &results)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(results), 5)
}
