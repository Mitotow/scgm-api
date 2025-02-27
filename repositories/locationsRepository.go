package repositories

import (
	"errors"
	"github.com/Mitotow/scgm-api/models"
	"gorm.io/gorm"
)

type LocationsRepository interface {
	FindAll(skip int, take int) ([]models.Location, int64, error)
	FindByName(name string) (models.Location, error)
}

type LocationsRepositoryImpl struct {
	Db *gorm.DB
}

func NewLocationsRepositoryImpl(Db *gorm.DB) LocationsRepository {
	return &LocationsRepositoryImpl{Db: Db}
}

func (r LocationsRepositoryImpl) FindAll(skip int, take int) ([]models.Location, int64, error) {
	var locations []models.Location
	var total int64

	res := r.Db.Model(&models.Location{}).Count(&total)
	if res.Error != nil {
		return nil, 0, res.Error
	}

	res = r.Db.Offset(skip).Limit(take).Find(&locations)
	if res.Error != nil {
		return nil, 0, res.Error
	}

	return locations, total, nil
}

func (r LocationsRepositoryImpl) FindByName(name string) (models.Location, error) {
	var location models.Location

	if err := r.Db.Where("name=?", name).First(&location).Error; err != nil {
		return location, errors.ErrUnsupported
	}

	return location, nil
}
