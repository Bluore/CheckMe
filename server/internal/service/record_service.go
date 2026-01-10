package service

import (
	"checkme/config"
	"checkme/internal/repository"
)

type RecordService interface {
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
