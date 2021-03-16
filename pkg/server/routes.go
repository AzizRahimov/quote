package app

func (s *Server) Init()  {
	s.router.POST("/quote", s.handleCreateQuote)
	s.router.GET("/quotes", s.handlerGetAllQuotes)
	s.router.POST("/editquote", s.handlerEditQuote)
	s.router.GET("/quotes/:category", s.handleGetQuoteByCategory)
	s.router.DELETE("/quote/:id", s.handleRemoveQuoteByID)


}
