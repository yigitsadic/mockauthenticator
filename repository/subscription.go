package repository

import "github.com/yigitsadic/mockauthenticator/models"

// SubscriptionRepository an interface for repository structure. Contains list and create functionalities.
type SubscriptionRepository interface {
	FindAllFor(owner string) ([]models.Subscription, error)
	Create(owner, issuedFor string) (*models.Subscription, error)
	CreateCode(owner, issuedFor string) (*models.Code, error)
}
