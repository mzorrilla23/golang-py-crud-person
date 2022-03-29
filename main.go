package main

import (
	"context"
	"fmt"
	"log"

	"example.com/sarang-apis/controllers"
	"example.com/sarang-apis/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server           *gin.Engine
	personservice    services.PersonService
	personcontroller controllers.PersonController
	ctx              context.Context
	personcollection *mongo.Collection
	mongoclient      *mongo.Client
	err              error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection established")

	personcollection = mongoclient.Database("persondb").Collection("persons")
	personservice = services.NewPersonService(personcollection, ctx)
	personcontroller = controllers.New(personservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/V1")
	personcontroller.RegisterPersonRouters(basepath)
	log.Fatal(server.Run(":9090"))
}