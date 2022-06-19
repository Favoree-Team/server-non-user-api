package repository

import (
	"github.com/Favoree-Team/server-non-user-api/entity"
	"gorm.io/gorm"
)

type PersonalDataRepository interface {
	GetByIDorIP(id string, ipAddress string) (entity.PersonalData, error)
	Create(personalData entity.PersonalData) error
	Update(id string, updateData map[string]interface{}) error
}

type personalDataRepository struct {
	db *gorm.DB
}

func NewPersonalDataRepository(db *gorm.DB) *personalDataRepository {
	return &personalDataRepository{db: db}
}

func (r *personalDataRepository) GetByIDorIP(id string, ipAddress string) (entity.PersonalData, error) {
	var personalData entity.PersonalData
	if id != "" && ipAddress != "" {
		if err := r.db.Where("id = ? OR ip_address = ?", id, ipAddress).First(&personalData).Error; err != nil {
			return personalData, err
		}
	} else if id != "" {
		if err := r.db.Where("id = ?", id).First(&personalData).Error; err != nil {
			return personalData, err
		}
	} else if ipAddress != "" {
		if err := r.db.Where("ip_address = ?", ipAddress).First(&personalData).Error; err != nil {
			return personalData, err
		}
	}
	return personalData, nil
}

func (r *personalDataRepository) Create(personalData entity.PersonalData) error {
	return r.db.Create(&personalData).Error
}

func (r *personalDataRepository) Update(id string, updateData map[string]interface{}) error {
	return r.db.Model(&entity.PersonalData{}).Where("id = ?", id).Updates(updateData).Error
}
