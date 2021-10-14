package repository

import "github.com/yigitsadic/mockauthenticator/models"

// InMemoryRepository implements repository interface with in-memory
type InMemoryRepository struct {
}

// Create mock create
func (m InMemoryRepository) Create(owner, issuedFor string) (*models.Subscription, error) {
	panic("implement me")
}

// CreateCode mock create code
func (m InMemoryRepository) CreateCode(owner, issuedFor string) (*models.Code, error) {
	panic("implement me")
}

// FindAllFor mock list subscriptions.
func (m InMemoryRepository) FindAllFor(owner string) ([]models.Subscription, error) {
	records := []models.Subscription{
		{
			ID:       1,
			Owner:    "3",
			IssuedTo: "e",
		},
		{
			ID:       2,
			Owner:    "4",
			IssuedTo: "5",
		},
	}

	return records, nil
}
