package controllers

import (
	"net/http"

	"github.com/amr9mohamed/mainApp/api/models"
	"github.com/amr9mohamed/mainApp/api/responses"
)

func (s *Server) getUsers(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.GetUsers(s.DB)
	if err != nil {
		responses.ERROR(rw, http.StatusInternalServerError, err)
	}
	responses.JSON(rw, http.StatusOK, users)
}
