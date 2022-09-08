package mongodb_repositories

import (
	"server/internal/app/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ApplicationsRepository struct {
	dbClient *mongo.Client
	logger   *zap.Logger
}

var _ repositories.ApplicationsRepository = (*ApplicationsRepository)(nil)

func NewApplicationsRepository(db *mongo.Client, l *zap.Logger) *ApplicationsRepository {
	return &ApplicationsRepository{
		dbClient: db,
		logger:   l,
	}
}

func (r *ApplicationsRepository) GetApplicationNames() ([]string, error) {
	return nil, nil
}
