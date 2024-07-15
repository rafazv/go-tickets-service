package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSpotNumber    = errors.New("invalid spot number")
	ErrSpotNotFound      = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "AVAILABLE"
	SpotStatusSold      SpotStatus = "SOLD"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string // all name uses the rule: Letter+Number. Ex: A1, B2, C3, etc.
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}
	if err := spot.Validate(); err != nil {
		return nil, err
	}
	return spot, nil
}

func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return errors.New("spot name is required")
	}
	if len(s.Name) < 2 {
		return errors.New("spot name must be at least 2 characters long")
	}
	// Validate if the spot name is in the correct format
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return errors.New("spot name must start with a letter")
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return errors.New("spot name must end with a number")
	}
	return nil
}