package services

import (
	"github.com/Mitotow/scgm-api/config"
	"github.com/Mitotow/scgm-api/models"
	"github.com/Mitotow/scgm-api/repositories"
	"net/http"
)

type LocationsService interface {
	FindAll(page int) (*models.LocationsResponse, *models.ErrorResponse)
	FindByName(name string) (*models.LocationResponse, *models.ErrorResponse)
}

type LocationsServiceImpl struct {
	LocationsRepository repositories.LocationsRepository
	env                 *config.EnvironmentVariables
}

func NewLocationsService(repository repositories.LocationsRepository) LocationsService {
	return &LocationsServiceImpl{LocationsRepository: repository, env: config.GetEnv()}
}

func (s LocationsServiceImpl) FindAll(page int) (*models.LocationsResponse, *models.ErrorResponse) {
	locations, total, err := s.LocationsRepository.FindAll(s.env.LocationsPerPage*(page-1), s.env.LocationsPerPage)
	if err != nil {
		return nil, &models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  "Internal Server Error",
		}
	}

	return &models.LocationsResponse{
		Status:     http.StatusOK,
		Page:       page,
		MaxPerPage: s.env.LocationsPerPage,
		Total:      total,
		Locations:  locations,
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
