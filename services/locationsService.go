package services

import "github.com/Mitotow/scgm-api/repositories"

type LocationsService interface {
	FindAll()
	FindByName()
}

type LocationsServiceImpl struct {
	LocationsRepository repositories.LocationsRepository
}

func NewLocationsService(repository repositories.LocationsRepository) LocationsService {
	return &LocationsServiceImpl{LocationsRepository: repository}
}

func (s LocationsServiceImpl) FindAll() {

}

func (s LocationsServiceImpl) FindByName() {

}
