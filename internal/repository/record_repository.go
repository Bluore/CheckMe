package repository

import (
	"checkme/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type RecordRepository interface {
	Create(ctx context.Context, record *model.Record) error
	GetByID(ctx context.Context, id string) (*model.Record, error)
	GetLastByDevice(ctx context.Context, device string) (*model.Record, error)
	GetDevice(ctx context.Context) (*[]string, error)
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
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (r *recordRepository) GetLastByDevice(ctx context.Context, device string) (*model.Record, error) {
	var record model.Record
	err := r.db.WithContext(ctx).Order("updated_time DESC").Where("device = ?", device).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (r *recordRepository) Update(ctx context.Context, record *model.Record) error {
	return r.db.WithContext(ctx).Save(record).Error
}

func (r *recordRepository) GetDevice(ctx context.Context) (*[]string, error) {
	var devices []string
	err := r.db.WithContext(ctx).Model(&model.Record{}).Distinct("device").Pluck("device", &devices).Error
	if err != nil {
		return nil, err
	}
	return &devices, nil
}
