package controllers

import (
	"net/http"
	"server/internal/app/services"

	"go.uber.org/zap"
)

type CoreDumpsController struct {
	service services.CoreDumpsService
	logger  *zap.Logger
}

func NewCoreDumpsController(s services.CoreDumpsService, l *zap.Logger) *CoreDumpsController {
	return &CoreDumpsController{
		service: s,
		logger:  l,
	}
}

func (c *CoreDumpsController) GetCoreDumps() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c *CoreDumpsController) AddCoreDump() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c *CoreDumpsController) DeleteCoreDump() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
