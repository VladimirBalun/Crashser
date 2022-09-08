package mongodb_repositories

import (
	"server/internal/app/entities"
	"server/internal/app/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type CoreDumpsRepository struct {
	dbClient *mongo.Client
	logger   *zap.Logger
}

var _ repositories.CoreDumpsRepository = (*CoreDumpsRepository)(nil)

func NewCoreDumpsRepository(db *mongo.Client, l *zap.Logger) *CoreDumpsRepository {
	return &CoreDumpsRepository{
		dbClient: db,
		logger:   l,
	}
}

func (r *CoreDumpsRepository) GetCoreDumps() ([]entities.CoreDump, error) {
	return nil, nil
}

func (r *CoreDumpsRepository) AddCoreDump(coreDump entities.CoreDump) error {
	return nil
}

func (r *CoreDumpsRepository) DeleteCoreDump(id string) error {
	return nil
}
