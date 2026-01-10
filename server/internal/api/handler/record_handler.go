package handler

import "checkme/internal/service"

type Handler struct {
	recordService service.RecoderService
}

func NewHandler(recoderService service.RecoderService) *Handler {
	return &Handler{recordService: recoderService}
}
