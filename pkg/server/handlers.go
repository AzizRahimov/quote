package app

import (
	"encoding/json"
	"github.com/AzizRahimov/quote/pkg/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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
	}

	err = s.quotes.CreateQuote(quote)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	resp, err := json.Marshal(quote)
	if err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err)
	}

}