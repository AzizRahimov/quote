package app

func (s *Server) Init()  {
	s.router.POST("/quote", s.handleCreateQuote) // Create a quote with following fields: author, quote, category.
	s.router.GET("/quotes", s.handlerGetAllQuotes) // Get all quotes.
	s.router.POST("/editquote", s.handlerEditQuote) // Edit/Change quote: author, quote, category.
	s.router.GET("/quotes/:category", s.handleGetQuoteByCategory) //  Get all quotes by category.
	s.router.DELETE("/quoteeee/:id", s.handleRemoveQuoteByID) // Delete a quote by ID.
	s.router.GET("/randomquote", s.handleGetRandomQuote) //  Get a random quote.


}
