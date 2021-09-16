package controllers

func (s *Server) initializeRoutes() {
	//Users routes
	s.Router.GET("/users", s.GetUsers)
	s.Router.GET("/users/:country/:pageNumber/:pageSize/", s.GetUsersByCountry)
	s.Router.GET("/users/countries", s.GetDistinctCountries)
	s.Router.GET("/users/frequency", s.GetCountryFrequency)
}
