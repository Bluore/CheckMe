package service

import (
	"checkme/config"
	"checkme/internal/repository"
)

type RecoderService interface {
}

type recoderService struct {
	recordRepo repository.UserRepository
	cfg        *config.Config
}

func NewRecoderService(recoderRepo repository.UserRepository, cfg *config.Config) RecoderService {
	return &recoderService{
		recordRepo: recoderRepo,
		cfg:        cfg,
	}
}
