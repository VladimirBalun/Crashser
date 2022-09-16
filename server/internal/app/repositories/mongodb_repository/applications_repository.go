package mongodb_repository

import (
	"context"
	"fmt"
	"server/internal/app/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ApplicationsRepository struct {
	ctx        context.Context
	dbClient   *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
}

var _ repositories.ApplicationsRepository = (*ApplicationsRepository)(nil)

func NewApplicationsRepository(ctx context.Context, db *mongo.Client, l *zap.Logger) *ApplicationsRepository {
	mongoDB := db.Database("crasher")
	collection := mongoDB.Collection("coredumps")

	return &ApplicationsRepository{
		ctx:        ctx,
		dbClient:   db,
		collection: collection,
		logger:     l,
	}
}

func (r *ApplicationsRepository) GetApplicationNames() ([]string, error) {
	timeout, err := ParseCtxTimeoutEnv()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	res, err := r.collection.Distinct(ctx, "appinfo.name", bson.D{})
	if err != nil {
		return nil, err
	}
	applications := make([]string, len(res))
	for i, v := range res {
		applications[i] = fmt.Sprint(v)
	}
	return applications, nil
}
