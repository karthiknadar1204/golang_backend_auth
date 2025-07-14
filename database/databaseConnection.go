package database

// This file is responsible for connecting to the MongoDB database and returning a client

import(
"fmt" // print messages
"log" // log messages
"time" // For timeout duration
"os" // To get environment variables
"context" //  Used to set a timeout for MongoDB operations
"github.com/joho/godotenv" // Loads variables from .env file
"go.mongodb.org/mongo-driver/mongo" // MongoDB driver
"go.mongodb.org/mongo-driver/mongo/options" // MongoDB options
)


// It’s like:
// DBinstance() → Connect to the database.
// OpenCollection() → Get a table (collection) from it.


// This function connects to MongoDB and returns a MongoDB client
// Return type is *mongo.Client (a pointer to a MongoDB connection object)
func DBinstance() *mongo.Client{

	//for loading environment variables from the .env file
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}

	//  Read the connection string
	MongoDb := os.Getenv("MONGODB_URL") //-> This is how you read environment variables in Go

	//This creates a new MongoDB client using the URL
	client, err:= mongo.NewClient(options.Client().ApplyURI(MongoDb)) //-> This is how you create a new MongoDB client
	if err != nil {
		log.Fatal(err)
	}

	// This sets a 10-second timeout while trying to connect
	// defer cancel() means “cancel the context after the function finishes”
	// client.Connect(ctx) actually opens the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
    
	//Return the connected client
	fmt.Println("Connected to MongoDB!")
	return client
}


// Global variable: Client
var Client *mongo.Client = DBinstance()
// This line runs DBinstance() once when the app starts
// It stores the connected client in a global variable Client
// Other files can use database.Client to access the DB



// Takes in a mongo.Client and a collection name (like "users" or "products")
// Connects to the "cluster0" database (change this to your DB name)
// Returns the collection object — a handle to perform find, insert, etc.
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
