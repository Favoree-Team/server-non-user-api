package repository

import "gorm.io/gorm"

type LogReqTransRepository interface {
}

type logReqTransRepository struct {
	db *gorm.DB
}

func NewLogTransRepository(db *gorm.DB) *logReqTransRepository {
	return &logReqTransRepository{db: db}
}
