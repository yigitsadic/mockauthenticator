package models

import "time"

// Code represents generated code for given owner and issuer.
type Code struct {
	Value     string    `json:"value"`
	Domain    string    `json:"domain"`
	ExpiresAt time.Time `json:"expires_at"`
}
