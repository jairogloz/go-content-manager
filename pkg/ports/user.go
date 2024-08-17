package ports

import "github.com/jairogloz/go-content-manager/pkg/domain"

type UserService interface {
	Auth(apiKey string) (*domain.User, error)
}

type UserRepository interface {
	FindByAPIKey(apiKey string) (*domain.User, error)
}
