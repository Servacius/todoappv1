package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string for localhhost mongoDB
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "test"

// Collection name
const collName = "todolist"

// Collection object/instance
var collection *mongo.Collection

// Create connection with mongo db
func init() {
	// Set client  options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}
