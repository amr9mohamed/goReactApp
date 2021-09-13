package controllers

import "github.com/amr9mohamed/mainApp/api/middlewares"

func (s *Server) initializeRoutes() {
	//Users routes
	s.Router.HandleFunc("/users", middlewares.JsonMiddleware(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{country}/{pageNumber}/{pageSize}/", middlewares.JsonMiddleware(s.GetUsersByCountry)).Methods("GET")
	s.Router.HandleFunc("/users/countries", middlewares.JsonMiddleware(s.GetDistinctCountries)).Methods("GET")
	s.Router.HandleFunc("/users/frequency", middlewares.JsonMiddleware(s.GetCountryFrequency)).Methods("GET")
}
