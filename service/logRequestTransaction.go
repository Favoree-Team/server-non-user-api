package service

import "github.com/Favoree-Team/server-non-user-api/repository"

type LogReqTransService interface {
}

type logReqTransService struct {
	logReqTransRepo repository.LogReqTransRepository
}

func NewLogReqTransService(logReqTransRepo repository.LogReqTransRepository) *logReqTransService {
	return &logReqTransService{logReqTransRepo: logReqTransRepo}
}
