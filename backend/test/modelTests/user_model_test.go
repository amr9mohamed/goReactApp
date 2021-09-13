package modeltests

import (
	"testing"

	"github.com/amr9mohamed/mainApp/api/controllers"
	"github.com/amr9mohamed/mainApp/api/models"
	"github.com/go-playground/assert/v2"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestGetUsers(t *testing.T) {
	users, err := userInstance.GetUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 1007616)
}

func TestGetUsersByCountry(t *testing.T) {
	users, err := userInstance.GetUsersByCountry(server.DB, "Cameroon", 1, 100)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 100)
}

func TestGetDistinctCountries(t *testing.T) {
	countries, err := userInstance.GetDistinctCountries(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*countries), 5)
}

func TestGetCountyFrequency(t *testing.T) {
	result, err := userInstance.GetDistinctCountries(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*result), 5)
}
