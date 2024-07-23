package usecase


type EventDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Organization string `json:"organization"`
	Rating       string `json:"rating"`
	Date         string `json:"date"`
	ImageURL     string `json:"image_url"`
	Capacity     int    `json:"capacity"`
	Price        int    `json:"price"`
	PartnerID    int    `json:"partner_id"`
}

type SpotDTO struct {
	ID       string `json:"id"`
	Name		 string `json:"name"`
	Status   string `json:"status"`
	TicketID string `json:"ticket_id"`
	EventID string `json:"event_id"`
	Reserved bool   `json:"reserved"`
}