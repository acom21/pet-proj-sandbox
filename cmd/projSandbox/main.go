package main

import (
	"context"
	"log"

	"github.com/acom21/pet-proj-dsandbox/config"
	"github.com/acom21/pet-proj-dsandbox/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	srv := service.NewService(cfg)

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
