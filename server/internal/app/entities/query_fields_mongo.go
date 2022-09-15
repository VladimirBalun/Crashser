package entities

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OptionsMongo func(bson.M, *options.FindOptions)

func SetApplication(application string) OptionsMongo {
	return func(filter bson.M, options *options.FindOptions) {
		filter["appinfo.name"] = application
	}
}

func SetStartTimestamp(start primitive.DateTime) OptionsMongo {
	return func(filter bson.M, options *options.FindOptions) {
		filter["timestamp"] = bson.M{
			"$gte": start,
		}
	}
}

func SetEndTimestamp(end primitive.DateTime) OptionsMongo {
	return func(filter bson.M, options *options.FindOptions) {
		filter["timestamp"] = bson.M{
			"$lt": end,
		}
	}
}

func SetLimit(limit int) OptionsMongo {
	return func(filter bson.M, options *options.FindOptions) {
		options.SetLimit(int64(limit))
	}
}

func SetOffset(offset int) OptionsMongo {
	return func(filter bson.M, options *options.FindOptions) {
		options.SetSkip(int64(offset))
	}
}
