package mongodb_repository

import (
	"context"
	"server/internal/app/entities"
	"server/internal/app/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type CoreDumpsRepository struct {
	ctx        context.Context
	dbClient   *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
}

var _ repositories.CoreDumpsRepository = (*CoreDumpsRepository)(nil)

func NewCoreDumpsRepository(ctx context.Context, db *mongo.Client, l *zap.Logger) *CoreDumpsRepository {
	mongoDB := db.Database("crasher")
	collection := mongoDB.Collection("coredumps")
	return &CoreDumpsRepository{
		ctx:        ctx,
		dbClient:   db,
		collection: collection,
		logger:     l,
	}
}

func (r *CoreDumpsRepository) GetCoreDumps(setters ...entities.OptionsMongo) ([]entities.CoreDump, error) {
	options := options.Find()
	filter := bson.M{}
	for _, setter := range setters {
		setter(filter, options)
	}
	var result []entities.CoreDump

	timeout, err := ParseCtxTimeoutEnv()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cur.All(r.ctx, &result)
	return result, err
}

func (r *CoreDumpsRepository) AddCoreDump(coreDump entities.CoreDump) error {
	timeout, err := ParseCtxTimeoutEnv()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	_, err = r.collection.InsertOne(ctx, coreDump)
	return err
}

func (r *CoreDumpsRepository) DeleteCoreDump(id string) error {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	timeout, err := ParseCtxTimeoutEnv()
	if err != nil {
		return  err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	_, err = r.collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: idHex}})
	return err
}

func (r *CoreDumpsRepository) DeleteAllCoreDumps() error {
	timeout, err := ParseCtxTimeoutEnv()
	if err != nil {
		return  err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	
	_, err = r.collection.DeleteMany(ctx, bson.D{})
	return err
}
