package repository

import "gorm.io/gorm"

type RecordRepository interface {
}

type recordRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}
