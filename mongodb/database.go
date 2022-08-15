package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB Atlas DB on Azure
const connectionString = "mongodb+srv://jongwon254:1234@cluster0.izncidl.mongodb.net/?retryWrites=true&w=majority"
const dbName = "chatroom"
const colName = "messages"

// db reference
var collection *mongo.Collection

func init() {
	// client
	clientOption := options.Client().ApplyURI(connectionString)

	// connection to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongodb connection successful")

	// collection ready
	collection = client.Database(dbName).Collection(colName)
}
