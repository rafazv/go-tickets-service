package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rafazv/go-tickets-service/internal/events/infra/repository"
	"github.com/rafazv/go-tickets-service/internal/events/infra/service"
	"github.com/rafazv/go-tickets-service/internal/events/usecase"

	httpHandler "github.com/rafazv/go-tickets-service/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(golang-mysql:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://host.docker.internal:8000/partner1",
		2: "http://host.docker.internal:8000/partner2",
	}
	
	partnerFactory := service.NewDefaultPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)
	createEventUseCase := usecase.NewCreateEventUseCase(eventRepo)
	createSpotsUseCase := usecase.NewCreateSpotsUseCase(eventRepo)

	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventUseCase,
		buyTicketsUseCase,
		createEventUseCase,
		createSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)
	r.HandleFunc("POST /events", eventsHandler.CreateEvent)
	r.HandleFunc("POST /events/{eventID}/spots", eventsHandler.CreateSpots)

	server := &http.Server{
		Addr:    ":7000",
		Handler: r,
	}

	log.Println("HTTP server running in port 7000")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error starting HTTP server: %v\n", err)
	}
}