package app

import (
	"encoding/json"
	"github.com/AzizRahimov/quote/pkg/models"
	"github.com/AzizRahimov/quote/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Server struct {
	 router  *httprouter.Router
	 quotes *models.Quotes
}

func NewServer(router *httprouter.Router, quotes *models.Quotes) *Server {
	return &Server{router: router, quotes: quotes}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	s.router.ServeHTTP(w, r)

}

func (s *Server) handleCreateQuote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	quote := models.Quote{}

	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	quote.CreatedAt = time.Now()
	err = s.quotes.CreateQuote(&quote)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	utils.SendResponse(w, quote)
}

func (s *Server) handlerGetAllQuotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	quotes, err := s.quotes.GetAllQuotes()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	utils.SendResponse(w, quotes)
}

func (s *Server) handlerEditQuote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	quote := &models.Quote{}
	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	editQuote, err := s.quotes.EditQuote(quote)
	if err != nil {
		log.Print(err)
		http.Error(w, "id not exist", http.StatusNotFound)
		return
	}
	utils.SendResponse(w, editQuote)
}

func (s *Server) handleRemoveQuoteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	quotes, err := s.quotes.Delete(id)
	if err == false {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	utils.SendResponse(w, quotes)
}

func (s *Server) handleGetQuoteByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := ps.ByName("category")

	quotes, err := s.quotes.GetQuotesByCategory(category)
	if err != nil {
		log.Print(err)
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	utils.SendResponse(w, quotes)

}

func (s *Server) handleGetRandomQuote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	quote, err := s.quotes.GetRandomQuote()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	utils.SendResponse(w, quote)

}