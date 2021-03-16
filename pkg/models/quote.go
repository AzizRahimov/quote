package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)


var ErrNotFound = errors.New("quotes not found")
var ErrIDNotFound = errors.New("ID quotes not found")

type Quote struct {
	ID string `json:"id"`
	Author string `json:"author"`
	Quote string `json:"quote"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"created_at"`

}

type Quotes struct {
	Quotes map[string]Quote
}

func NewQuotes() *Quotes {
	return &Quotes{Quotes: make(map[string]Quote)}
}



//Create Quotes
func (q *Quotes) CreateQuote(quote *Quote) (err error ) {
		fmt.Println(quote)

		quote.ID = uuid.New().String()
		q.Quotes[quote.ID] = *quote

		if q.Quotes == nil{
			return err
		}


		return   nil
	
}

func (q *Quotes) GetAll() ([]Quote, error) {
	quotes := []Quote{}
	for _, value := range q.Quotes{
		quotes = append(quotes, value)

	}
	if quotes == nil{
		return nil,  ErrNotFound
	}

	return quotes, nil
}

func (q *Quotes) EditQuote(quote *Quote) (*Quote, error)  {



		for key, _ := range q.Quotes{
			if key == quote.ID{
				q.Quotes[quote.ID] = *quote
				return quote, nil
			}

		}

	return nil, ErrIDNotFound

}

func (q *Quotes) DeleteQuoteByID(id string) ([]Quote, bool) {
	if id == ""{
		return nil, false
	}
	if _, ok := q.Quotes[id]; ok {
		delete(q.Quotes, id)
		quotes, _ := q.GetAll()
		return quotes, true
	}

	return nil, false
}


