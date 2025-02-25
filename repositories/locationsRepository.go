package repositories

import (
	"errors"
	"github.com/Mitotow/scgm-api/models"
	"gorm.io/gorm"
)

type LocationsRepository interface {
	FindAll() ([]models.Location, error)
	FindByName(name string) (models.Location, error)
}

type LocationsRepositoryImpl struct {
	Db *gorm.DB
}

func NewLocationsRepositoryImpl(Db *gorm.DB) LocationsRepository {
	return &LocationsRepositoryImpl{Db: Db}
}

func (r LocationsRepositoryImpl) FindAll() ([]models.Location, error) {
	var locations []models.Location
	res := r.Db.Find(&locations)

	return locations, res.Error
}

func (r LocationsRepositoryImpl) FindByName(name string) (models.Location, error) {
	var location models.Location

	if err := r.Db.Where("name=?", name).First(&location).Error; err != nil {
		return location, errors.New("location not found")
	}

	return location, nil
}
