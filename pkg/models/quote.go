package models

import (
	"github.com/google/uuid"
	"time"
)




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

		quote.ID = uuid.New().String()
		q.Quotes[quote.ID] = quote

		if q.Quotes == nil{
			return err
		}


		return   nil
	
}

