package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/jongwon254/Chat-Room/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Helper Methods

// insert message into database
func insertMessage(message model.Message) {
	inserted, err := collection.InsertOne(context.Background(), message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted 1 message with id: ", inserted.InsertedID)
}

// delete all records from db
func deleteAll() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleted.DeletedCount, "deleted")
	return deleted.DeletedCount
}

// get all records from db
func getAll() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var messages []primitive.M

	// loop through messages
	for cursor.Next(context.Background()) {
		var message bson.M
		err := cursor.Decode(&message)

		if err != nil {
			log.Fatal(err)
		}
		messages = append(messages, message)
	}

	defer cursor.Close(context.Background())
	return messages
}
