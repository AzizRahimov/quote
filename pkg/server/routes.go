package app

func (s *Server) Init()  {
	s.router.POST("/quote", s.handleCreateQuote)
	s.router.GET("/quotes", s.handlerGetAllQuotes)

}
