package content_item

import (
	"fmt"
	"math"
	"strings"

	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (s *Service) List(userID string, page, limit int, sortBy string, category string) (response *domain.ContentItemListResponse, err error) {

	if page < 1 {
		return nil, fmt.Errorf("error listing content items: page must be greater than or equal to 1")
	}

	if limit < 1 || limit > 100 {
		return nil, fmt.Errorf("error listing content items: limit must be between 1 and 100")
	}

	validSortByFileds := []string{"created_at", "updated_at", "category"}

	sortByOrder := "desc"
	sortByField := "updated_at"
	if sortBy != "" {
		// Break sortBy string into fields by "."
		sortByFields := strings.Split(sortBy, ".")
		if len(sortByFields) > 2 {
			return nil, fmt.Errorf("error listing content items: invalid sortBy field")
		}
		if len(sortByFields) == 2 {
			if sortByFields[1] != "asc" && sortByFields[1] != "desc" {
				return nil, fmt.Errorf("error listing content items: invalid sortBy field")
			}
			sortByOrder = sortByFields[1]
		}
		sortByField = sortByFields[0]
		if !contains(validSortByFileds, sortByField) {
			return nil, fmt.Errorf("error listing content items: invalid sortBy field")
		}
	}

	totalCount, err := s.Repo.Count(userID, category)
	if err != nil {
		return nil, err
	}

	// Calcualr número total de páginas
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	if page > totalPages {
		page = totalPages
	}

	contentItems, err := s.Repo.List(userID, page, limit, sortByField, sortByOrder, category)
	if err != nil {
		return nil, err
	}

	response = &domain.ContentItemListResponse{
		ContentItems: contentItems,
		TotalCount:   totalCount,
		Count:        len(contentItems),
		Page:         page,
		TotalPages:   totalPages,
	}

	return response, nil

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
