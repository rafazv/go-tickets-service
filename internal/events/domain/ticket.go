package domain

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