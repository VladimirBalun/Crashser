package mongodb_repository

import (
	"context"
	"server/internal/app"
	"server/internal/app/entities"
	"server/internal/app/repositories"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type CoreDumpsRepository struct {
	ctx        context.Context
	dbClient   *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
}

var _ repositories.CoreDumpsRepository = (*CoreDumpsRepository)(nil)

func NewCoreDumpsRepository(db *mongo.Client, l *zap.Logger) *CoreDumpsRepository {
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

	return &CoreDumpsRepository{
		ctx:        ctx,
		dbClient:   db,
		collection: collection,
		logger:     l,
	}
}

func (r *CoreDumpsRepository) GetCoreDumps() ([]entities.CoreDump, error) {
	var result []entities.CoreDump
	cur, err := r.collection.Find(r.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cur.All(r.ctx, &result)
	return result, err
}

func (r *CoreDumpsRepository) AddCoreDump(coreDump entities.CoreDump) error {
	_, err := r.collection.InsertOne(r.ctx, coreDump)
	return err
}

func (r *CoreDumpsRepository) DeleteCoreDump(id string) error {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(r.ctx, bson.D{{Key: "_id", Value: idHex}})
	return err
}
