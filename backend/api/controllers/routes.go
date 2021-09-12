package controllers

import "github.com/amr9mohamed/mainApp/api/middlewares"

func (s *Server) initializeRoutes() {
	//Users routes
	s.Router.HandleFunc("/users", middlewares.JsonMiddleware(s.getUsers)).Methods("GET")
	s.Router.HandleFunc("/users/countries", middlewares.JsonMiddleware(s.getDistinctCountries)).Methods("GET")
	s.Router.HandleFunc("/users/frequency", middlewares.JsonMiddleware(s.GetCountyFrequency)).Methods("GET")
}
