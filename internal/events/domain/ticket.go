package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

type TicketKind string

const (
	TicketKindHalf TicketKind = "HALF"
	TicketKindFull TicketKind = "FULL"
)

var ErrInvalidTicketPriceZero = errors.New("ticket price must be greater than 0")
var ErrInvalidTicketKind = errors.New("invalid ticket kind")

func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !IsValidTicketKind(ticketKind) {
		return nil, ErrInvalidTicketKind
	}

	t := &Ticket{
		ID: uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketKind,
		Price:      event.Price,
	}
	
	t.CalculatePrice()
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}

func IsValidTicketKind(TicketKind TicketKind) bool {
	return TicketKind == TicketKindHalf || TicketKind == TicketKindFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrInvalidTicketPriceZero
	}
	return nil
}