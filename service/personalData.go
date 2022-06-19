package service

import (
	"errors"
	"time"

	"github.com/Favoree-Team/server-non-user-api/auth"
	"github.com/Favoree-Team/server-non-user-api/entity"
	"github.com/Favoree-Team/server-non-user-api/helper"
	"github.com/Favoree-Team/server-non-user-api/repository"
)

type PersonalDataService interface {
	ValidationPersonalData(id string, ipAddress string) (entity.ValidationResponse, error)
	CreatePersonalData(data entity.CreatePersonalData) (entity.PersonalDataResponse, error)
	SubscribePersonalData(id string, ipAddress string, token string, data entity.CreateDetailPersonalData) (entity.PersonalDataResponse, error)
}

type personalDataService struct {
	personalDataRepository repository.PersonalDataRepository
	authService            auth.AuthService
}

func NewPersonalDataService(personalDataRepository repository.PersonalDataRepository, authService auth.AuthService) *personalDataService {
	return &personalDataService{
		personalDataRepository: personalDataRepository,
		authService:            authService,
	}
}

func (s *personalDataService) ValidationPersonalData(id string, ipAddress string) (entity.ValidationResponse, error) {
	personalData, err := s.personalDataRepository.GetByIDorIP(id, ipAddress)
	if err != nil {
		return entity.ValidationResponse{}, err
	}

	if personalData.Id == "" || personalData.IpAddress == "" {
		return entity.ValidationResponse{
			Validation:    entity.ValidationInvalid,
			MessageDetail: "ID or IP Address not found",
		}, nil
	}

	if ipAddress != personalData.IpAddress {
		return entity.ValidationResponse{
			Validation:    entity.ValidationInvalid,
			MessageDetail: "IP Address not match",
		}, nil
	}

	return entity.ValidationResponse{
		Validation:    entity.ValidationValid,
		MessageDetail: "success",
	}, nil
}

func (s *personalDataService) CreatePersonalData(data entity.CreatePersonalData) (entity.PersonalDataResponse, error) {
	// check the ip address
	personalData, err := s.personalDataRepository.GetByIDorIP("", data.IpAddress)
	if err != nil {
		return entity.PersonalDataResponse{}, err
	}

	if personalData.Id != "" && personalData.IpAddress != "" {
		generateToken, err := s.authService.GenerateToken(personalData.Id, personalData.IpAddress, personalData.DeviceAccess)
		if err != nil {
			return entity.PersonalDataResponse{}, err
		}

		return personalData.ToPersonalDataResponse(generateToken), nil
	}

	// create new personal data
	newPersonalData := entity.PersonalData{
		Id:           helper.NewUUID(),
		IpAddress:    data.IpAddress,
		DeviceAccess: data.DeviceAccess,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}

	err = s.personalDataRepository.Create(newPersonalData)
	if err != nil {
		return entity.PersonalDataResponse{}, err
	}

	generateNewToken, err := s.authService.GenerateToken(newPersonalData.Id, newPersonalData.IpAddress, newPersonalData.DeviceAccess)
	if err != nil {
		return entity.PersonalDataResponse{}, err
	}

	return newPersonalData.ToPersonalDataResponse(generateNewToken), nil
}

func (s *personalDataService) SubscribePersonalData(id string, ipAddress string, token string, data entity.CreateDetailPersonalData) (entity.PersonalDataResponse, error) {
	personalData, err := s.personalDataRepository.GetByIDorIP(id, ipAddress)
	if err != nil {
		return entity.PersonalDataResponse{}, err
	}

	if personalData.Id == "" || personalData.IpAddress == "" {
		return entity.PersonalDataResponse{}, errors.New("error: id or IP Address not found")
	}

	if ipAddress != personalData.IpAddress {
		return entity.PersonalDataResponse{}, errors.New("error: ip Address not match")
	}

	var updateData = make(map[string]interface{})

	updateData["updated_at"] = time.Now()

	if data.Name != "" {
		updateData["name"] = data.Name
		personalData.Name = data.Name
	}

	if data.Email != "" {
		updateData["email"] = data.Email
		personalData.Email = data.Email
	}

	err = s.personalDataRepository.Update(personalData.Id, updateData)
	if err != nil {
		return entity.PersonalDataResponse{}, err
	}

	return personalData.ToPersonalDataResponse(token), nil
}
