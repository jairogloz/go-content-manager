package user

import "github.com/jairogloz/go-content-manager/pkg/domain"

func (s *Service) Auth(apiKey string) (*domain.User, error) {
	user, err := s.Repo.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return user, nil
}
