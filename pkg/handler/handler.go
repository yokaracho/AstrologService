package handler

import "AstrologService/pkg/service"

type Handler struct {
	service service.Implementation
}

func NewHandler(service service.Implementation) *Handler {
	return &Handler{
		service: service,
	}
}
