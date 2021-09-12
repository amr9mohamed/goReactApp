package controllers

import (
	"net/http"

	"github.com/amr9mohamed/mainApp/api/models"
	"github.com/amr9mohamed/mainApp/api/responses"
	"github.com/gorilla/mux"
)

func (s *Server) getUsers(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.GetUsers(s.DB)
	if err != nil {
		responses.ERROR(rw, http.StatusInternalServerError, err)
	}
	responses.JSON(rw, http.StatusOK, users)
}

func (s *Server) getUsersByCountry(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if country, ok := mux.Vars(r)["country"]; ok {
		users, err := user.GetUsersByCountry(s.DB, country)
		if err != nil {
			responses.ERROR(rw, http.StatusInternalServerError, err)
		}
		responses.JSON(rw, http.StatusOK, users)
	} else {
		responses.ERROR(rw, http.StatusInternalServerError, nil)
	}
}

func (s *Server) getDistinctCountries(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	distinctCountries, err := user.GetDistinctCountries(s.DB)
	if err != nil {
		responses.ERROR(rw, http.StatusInternalServerError, err)
	}
	responses.JSON(rw, http.StatusOK, distinctCountries)
}

func (s *Server) GetCountyFrequency(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	countryFrequency, err := user.GetCountyFrequency(s.DB)
	if err != nil {
		responses.ERROR(rw, http.StatusInternalServerError, err)
	}
	responses.JSON(rw, http.StatusOK, countryFrequency)
}
