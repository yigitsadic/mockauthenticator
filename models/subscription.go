package models

// Subscription struct represents data in storage.
type Subscription struct {
	ID     int    `json:"id"`
	Owner  string `json:"owner"`
	Domain string `json:"domain"`
}
