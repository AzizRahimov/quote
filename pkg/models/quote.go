package models

import (
	"errors"
	"github.com/AzizRahimov/quote/utils"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)


var ErrNotFound = errors.New("quotes not found")
var ErrIDNotFound = errors.New("ID quotes not found")
var ErrMustBePositive = errors.New("number can't be zero")

type Quote struct {
	ID string `json:"id"`
	Author string `json:"author"`
	Quote string `json:"quote"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"created_at"`

}

type Quotes struct {
	Quotes map[string]Quote
	sync.Mutex
}

func NewQuotes() *Quotes {
	return &Quotes{Quotes: make(map[string]Quote)}
}



//Create Quotes
func (q *Quotes) CreateQuote(quote *Quote) (err error) {
	q.Lock()
	defer  q.Unlock()
	quote.ID = uuid.New().String()
	q.Quotes[quote.ID] = *quote

	if q.Quotes == nil {
		return err
	}
	return nil
}

// Get All Quotes
func (q *Quotes) GetAllQuotes() ([]Quote, error) {
	q.Lock()
	defer  q.Unlock()
	quotes := []Quote{}

	for _, value := range q.Quotes {
		quotes = append(quotes, value)

	}
	if quotes == nil {
		return nil, ErrNotFound
	}
	return quotes, nil
}

// EditQuote - edit quote by id
func (q *Quotes) EditQuote(quote *Quote) (*Quote, error) {
	q.Lock()
	defer  q.Unlock()

	for key, _ := range q.Quotes {
		if key == quote.ID {
			q.Quotes[quote.ID] = *quote
			return quote, nil
		}

	}
	return nil, ErrIDNotFound
}

func (q *Quotes) Delete(quoteID string) ([]Quote, bool) {

	_, exists := q.Quotes[quoteID]
	if exists {
		delete(q.Quotes, quoteID)
		quotes, _ := q.GetAllQuotes()
		return quotes, true
	}
	return nil, false
}





// Get Quotes by Category
func (q *Quotes) GetQuotesByCategory(category string) ([]Quote, error) {
	q.Lock()
	defer  q.Unlock()
	quotes := []Quote{}

	for _, value := range q.Quotes {
		if value.Category == category {
			quotes = append(quotes, value)

		}
	}
	if quotes == nil {
		return nil, ErrNotFound
	}

	return quotes, nil
}


// Get Random Quote
func (q *Quotes) GetRandomQuote() (*Quote, error) {
	q.Lock()
	defer  q.Unlock()
	rand.Seed(time.Now().UnixNano())
	count := 0
	randomNumber := rand.Intn(len(q.Quotes))
	if randomNumber == 0 {
		return nil, ErrMustBePositive
	}

	for _, quote := range q.Quotes {
		count++
		if count == randomNumber {
			return &quote, nil
		}

	}
	return nil, ErrNotFound
}


// Delete Quotes old quotes that were created 1 hour ago
func (q *Quotes) DeleteOldQuotes() {

	for _, quote := range q.Quotes {
		if utils.IsTimePassed(time.Now().Add( - time.Hour), quote.CreatedAt) {
			q.Delete(quote.ID)
		}
	}
}