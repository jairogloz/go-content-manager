package content_item

import "fmt"

func (s *Service) Delete(id string) error {
	if err := s.Repo.Delete(id); err != nil {
		return fmt.Errorf("error deleting content item: %w", err)
	}

	return nil
}
