package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = connectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://ricardo:<password>@gopher-twitter.rlkx5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func connectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexión exitosa con la BD")
	return client

}
