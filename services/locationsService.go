package services

import (
	"github.com/Mitotow/scgm-api/models"
	"github.com/Mitotow/scgm-api/repositories"
	"net/http"
)

type LocationsService interface {
	FindAll() (*models.LocationsResponse, *models.ErrorResponse)
	FindByName(name string) (*models.LocationResponse, *models.ErrorResponse)
}

type LocationsServiceImpl struct {
	LocationsRepository repositories.LocationsRepository
}

func NewLocationsService(repository repositories.LocationsRepository) LocationsService {
	return &LocationsServiceImpl{LocationsRepository: repository}
}

func (s LocationsServiceImpl) FindAll() (*models.LocationsResponse, *models.ErrorResponse) {
	locations, err := s.LocationsRepository.FindAll()
	if err != nil {
		return nil, &models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  "Internal Server Error",
		}
	}

	return &models.LocationsResponse{
		Status:    http.StatusOK,
		Total:     len(locations),
		Locations: locations,
	}, nil
}

func (s LocationsServiceImpl) FindByName(name string) (*models.LocationResponse, *models.ErrorResponse) {
	location, err := s.LocationsRepository.FindByName(name)
	if err != nil {
		return nil, &models.ErrorResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}
	}

	return &models.LocationResponse{
		Status:   http.StatusOK,
		Location: location,
	}, nil
}
