package handler

import "github.com/amanakin/k8s_hw/internal/service"

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type Handler struct {
	service service.Service
}
