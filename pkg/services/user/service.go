package user

import (
	"fmt"

	"github.com/jairogloz/go-content-manager/pkg/ports"
)

var _ ports.UserService = &Service{}

type Service struct {
	Repo ports.UserRepository
}

func NewService(repo ports.UserRepository) (*Service, error) {
	if repo == nil {
		return nil, fmt.Errorf("error creating user service: repository is nil")
	}

	return &Service{
		Repo: repo,
	}, nil
}
