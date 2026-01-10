package service

import (
	"checkme/config"
	"checkme/internal/dto"
	"checkme/internal/repository"
	"context"
)

type RecordService interface {
	Update(ctx context.Context, req *dto.UploadRecordRequest) error
}

type recordService struct {
	recordRepo repository.RecordRepository
	cfg        *config.Config
}

func NewRecoderService(recoderRepo repository.RecordRepository, cfg *config.Config) RecordService {
	return &recordService{
		recordRepo: recoderRepo,
		cfg:        cfg,
	}
}

func (rs recordService) Update(ctx context.Context, req *dto.UploadRecordRequest) error {

	return nil
}
