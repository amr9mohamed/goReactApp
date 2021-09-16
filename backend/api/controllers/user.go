package controllers

import (
	"net/http"
	"strconv"

	"github.com/amr9mohamed/mainApp/api/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetUsers(c *gin.Context) {
	user := models.User{}
	users, err := user.GetUsers(s.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (s *Server) GetUsersByCountry(c *gin.Context) {
	user := models.User{}
	pageNumber, err := strconv.Atoi(c.Param("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(c.Param("pageSize"))
	if err != nil {
		pageSize = 100
	}
	if country := c.Param("country"); country != "" {
		users, err := user.GetUsersByCountry(s.DB, country, pageNumber, pageSize)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(http.StatusOK, users)
	} else {
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}
}

func (s *Server) GetDistinctCountries(c *gin.Context) {
	user := models.User{}
	distinctCountries, err := user.GetDistinctCountries(s.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, distinctCountries)
}

func (s *Server) GetCountryFrequency(c *gin.Context) {
	user := models.User{}
	countryFrequency, err := user.GetCountyFrequency(s.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, countryFrequency)
}
