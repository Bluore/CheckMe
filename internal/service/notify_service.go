package service

import (
	"checkme/config"
	"checkme/internal/dto"
	"checkme/pkg/request"

	"github.com/gin-gonic/gin"
)

type NotifyService interface {
	CreateNotify(c *gin.Context, notify *dto.CreateNotifyRequest) error
}

type notifyService struct {
	cfg *config.Config
}

func NewNotifyService(cfg *config.Config) NotifyService {
	return &notifyService{
		cfg: cfg,
	}
}

func (ns *notifyService) CreateNotify(c *gin.Context, notify *dto.CreateNotifyRequest) error {
	go request.NotifyFeishu(ns.cfg, c, *notify)

	return nil
}
