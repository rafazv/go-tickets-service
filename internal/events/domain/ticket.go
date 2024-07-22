package domain

import "errors"

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

type TicketType string

const (
	TicketTypeHalf TicketType = "HALF"
	TicketTypeFull TicketType = "FULL"
)

var ErrInvalidTicketPriceZero = errors.New("ticket price must be greater than 0")

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrInvalidTicketPriceZero
	}
	return nil
}