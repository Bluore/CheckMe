package repository

import (
	"checkme/internal/model"
	"context"

	"gorm.io/gorm"
)

type RecordRepository interface {
	Create(ctx context.Context, record *model.Record) error
	GetByID(ctx context.Context, id string) (*model.Record, error)
	GetLastByDevice(ctx context.Context, device string) (*model.Record, error)
	Update(ctx context.Context, record *model.Record) error
}

type recordRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}

func (r *recordRepository) Create(ctx context.Context, record *model.Record) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *recordRepository) GetByID(ctx context.Context, id string) (*model.Record, error) {
	var record model.Record
	err := r.db.WithContext(ctx).Where("id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (r *recordRepository) GetLastByDevice(ctx context.Context, device string) (*model.Record, error) {
	var record model.Record
	err := r.db.WithContext(ctx).Order("updated_time DESC").Where("device = ?", device).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (r *recordRepository) Update(ctx context.Context, record *model.Record) error {
	return r.db.WithContext(ctx).Save(record).Error
}
