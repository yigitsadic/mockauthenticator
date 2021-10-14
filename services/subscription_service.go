package services

import (
	"github.com/yigitsadic/mockauthenticator/models"
	"github.com/yigitsadic/mockauthenticator/repository"
)

// SubscriptionService handles interactions with given storage.
type SubscriptionService struct {
	repository repository.SubscriptionRepository
}

// NewSubscriptionService initializes new service with given repository.
func NewSubscriptionService(repository repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repository: repository}
}

// FindAll fetches all records from storage.
func (s *SubscriptionService) FindAll(owner string) ([]models.Subscription, error) {
	result, err := s.repository.FindAllFor(owner)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Create creates new record from service storage with given parameters.
func (s *SubscriptionService) Create(owner, issuedTo string) (*models.Subscription, error) {
	return nil, nil
}
