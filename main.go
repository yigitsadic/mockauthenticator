package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yigitsadic/mockauthenticator/handlers"
	"github.com/yigitsadic/mockauthenticator/repository"
	"github.com/yigitsadic/mockauthenticator/services"
	"log"
	"net/http"
	"os"
)

func main() {
	port, portPresent := os.LookupEnv("PORT")
	if !portPresent {
		port = "5050"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5, "application/json"))
	r.Use(middleware.Heartbeat("/health"))

	service := services.NewSubscriptionService(new(repository.InMemoryRepository))

	handler := handlers.SubscriptionHandler{SubscriptionService: service}

	r.Get("/subscriptions", handler.HandleListSubscriptions)
	r.Post("/subscriptions", handler.HandleCreateSubscription)
	r.Post("/codes", handler.HandleCreateCode)

	log.Println("Server is up and running on " + port)
	panic(http.ListenAndServe(":"+port, r))
}
