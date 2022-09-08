package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/app/services"

	"go.uber.org/zap"
)

type ApplicationsController struct {
	service services.ApplicationsService
	logger  *zap.Logger
}

func NewApplicationsController(s services.ApplicationsService, l *zap.Logger) *ApplicationsController {
	return &ApplicationsController{
		service: s,
		logger:  l,
	}
}

func (c *ApplicationsController) GetApplicationNames() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		appNames, err := c.service.GetApplicationNames()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			c.logger.Error("get application names failed", zap.Error(err))
			return
		}

		response, err := json.Marshal(appNames)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			c.logger.Error("response marshaling failed", zap.Error(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}
}
