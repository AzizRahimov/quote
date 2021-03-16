package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)


var ErrNotFound = errors.New("quotes not found")

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



//func NewQuotes(quotes map[string]Quote) *Quotes {
//	return &Quotes{Quotes: quotes}
//}

//Create Quotes
func (q *Quotes) CreateQuote(quote Quote) (err error ) {
		fmt.Println(quote)

		quote.ID = uuid.New().String()
		q.Quotes[quote.ID] = quote

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

