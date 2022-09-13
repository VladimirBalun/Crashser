package mongodb_repository

import (
	"context"
	"fmt"
	"server/internal/app"
	"server/internal/app/repositories"
	"strconv"
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

func NewApplicationsRepository(db *mongo.Client, l *zap.Logger) *ApplicationsRepository {
	mongoDB := db.Database("crasher")
	collection := mongoDB.Collection("coredumps")

	ctxTimeout, err := strconv.Atoi(app.CtxTimeout)
	if err != nil {
		l.Fatal(
			"failed to convert context timeout to number",
			zap.Error(err),
		)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Second)
	defer cancel()

	return &ApplicationsRepository{
		ctx:        ctx,
		dbClient:   db,
		collection: collection,
		logger:     l,
	}
}

func (r *ApplicationsRepository) GetApplicationNames() ([]string, error) {
	res, err := r.collection.Distinct(r.ctx, "app_info.name", bson.D{})
	if err != nil {
		return nil, err
	}
	// convert []interface to []string
	strArr := make([]string, len(res))
	for i, v := range res {
		strArr[i] = fmt.Sprint(v)
	}
	return strArr, nil
}
