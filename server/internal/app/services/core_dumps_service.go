package services

import (
	"server/internal/app/entities"
	"server/internal/app/repositories"

	"go.uber.org/zap"
)

type CoreDumpsService interface {
	GetCoreDumps() ([]entities.CoreDump, error)
	AddCoreDump(coreDump entities.CoreDump) error
	DeleteCoreDump(id string) error
}

type CoreDumpServiceImpl struct {
	repository repositories.CoreDumpsRepository
	logger     *zap.Logger
}

func NewCoreDumpsService(r repositories.CoreDumpsRepository, l *zap.Logger) CoreDumpsService {
	return &CoreDumpServiceImpl{
		repository: r,
		logger:     l,
	}
}

func (s *CoreDumpServiceImpl) GetCoreDumps() ([]entities.CoreDump, error) {
	return nil, nil
}

func (s *CoreDumpServiceImpl) AddCoreDump(coreDump entities.CoreDump) error {
	return nil
}

func (s *CoreDumpServiceImpl) DeleteCoreDump(id string) error {
	return nil
}
