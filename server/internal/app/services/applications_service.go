package services

import (
	"server/internal/app/repositories"

	"go.uber.org/zap"
)

type ApplicationsService interface {
	GetApplicationNames() ([]string, error)
}

type ApplicationsServiceImpl struct {
	repository repositories.ApplicationsRepository
	logger     *zap.Logger
}

func NewApplicationsService(r repositories.ApplicationsRepository, l *zap.Logger) ApplicationsService {
	return &ApplicationsServiceImpl{
		repository: r,
		logger:     l,
	}
}

func (s *ApplicationsServiceImpl) GetApplicationNames() ([]string, error) {
	return nil, nil
}
