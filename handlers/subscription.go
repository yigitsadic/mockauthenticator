package handlers

import (
	"encoding/json"
	"github.com/yigitsadic/mockauthenticator/dtos"
	"github.com/yigitsadic/mockauthenticator/services"
	"net/http"
)

// SubscriptionHandler stores subscription service for its handler functions.
type SubscriptionHandler struct {
	SubscriptionService *services.SubscriptionService
}

// HandleListSubscriptions lists subscriptions for given user email in header.
func (h *SubscriptionHandler) HandleListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := h.SubscriptionService.FindAll("lorem")
	if err != nil {
		json.NewEncoder(w).Encode(dtos.GeneralErrorResponse{Message: "Hello, error occurred"})
		return
	}

	json.NewEncoder(w).Encode(result)
}

// HandleCreateSubscription creates subscription for given domain and owner.
func (h *SubscriptionHandler) HandleCreateSubscription(w http.ResponseWriter, r *http.Request) {

}

// HandleCreateCode creates codes for given domain and owner.
func (h *SubscriptionHandler) HandleCreateCode(w http.ResponseWriter, r *http.Request) {

}
