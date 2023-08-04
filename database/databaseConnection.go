package database

import (
	"context"
	"fmt"
	"log"

	//"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading the .env file")
	// }

	//MongoDb := os.Getenv("MONGODB_URL")
	fmt.Println("creating mongo db instance!!")
	const MongoDb string = "mongodb+srv://ysrishu1031:aabRvg0MevdpuUYK@cluster0.yzrcqxc.mongodb.net/?retryWrites=true&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		fmt.Println("error in mongo connection!!1")
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!!")

	return client
}

var CLient *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
