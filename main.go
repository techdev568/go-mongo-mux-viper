package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/parsaakbari1209/go-mongo-crud-rest-api/config"
	"github.com/parsaakbari1209/go-mongo-crud-rest-api/db"
	"github.com/parsaakbari1209/go-mongo-crud-rest-api/repository"
	"github.com/parsaakbari1209/go-mongo-crud-rest-api/routes"
)

func main() {
	log1 := log.New()
	log1.Formatter = new(log.JSONFormatter)
	log.Info("To do App Initiated")

	config, err := config.LoadConfig("./config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Info("Port:", config.Port)

	client, err := db.ConnectToDB(config.MongodbURL)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// create a repository
	repository := repository.NewRepository(client.Database(config.Database))

	router := routes.Routes(repository)
	log1.Fatal(http.ListenAndServe(":"+config.Port, router))
}
