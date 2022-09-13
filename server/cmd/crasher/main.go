package main

import (
	"context"
	"net/http"
	"server/internal/app"
	"server/internal/app/controllers"
	"server/internal/app/repositories/mongodb_repository"
	"server/internal/app/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(app.DatabaseAddress))
	if err != nil {
		logger.Fatal(
			"failed to connect with database",
			zap.String("address", app.DatabaseAddress),
			zap.Error(err),
		)
	}

	if err := dbClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		logger.Fatal(
			"failed to ping database",
			zap.String("address", app.DatabaseAddress),
			zap.Error(err),
		)
	}

	defer func() {
		if err = dbClient.Disconnect(context.TODO()); err != nil {
			logger.Error(
				"failed to disconnect with database",
				zap.String("address", app.DatabaseAddress),
				zap.Error(err),
			)
		}
	}()
	appsRepo := mongodb_repository.NewApplicationsRepository(dbClient, logger)
	coreDumpsRepo := mongodb_repository.NewCoreDumpsRepository(dbClient, logger)

	appsService := services.NewApplicationsService(appsRepo, logger)
	coreDumpsService := services.NewCoreDumpsService(coreDumpsRepo, logger)

	appsController := controllers.NewApplicationsController(appsService, logger)
	coreDumpsController := controllers.NewCoreDumpsController(coreDumpsService, logger)

	router := mux.NewRouter()
	router.HandleFunc("/core_dumps", coreDumpsController.GetCoreDumps()).Methods("GET")
	router.HandleFunc("/core_dumps", coreDumpsController.AddCoreDump()).Methods("POST")
	router.HandleFunc("/core_dumps/{id}", coreDumpsController.DeleteCoreDump()).Methods("DELETE")
	router.HandleFunc("/applications", appsController.GetApplicationNames()).Methods("GET")

	err = http.ListenAndServe(app.ServerAddress, router)
	if err != nil {
		logger.Fatal(
			"failed to start server",
			zap.String("address", app.ServerAddress),
			zap.Error(err),
		)
	}
}
