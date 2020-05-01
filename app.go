package main

import (
	// db "./app/common/libs/db.go"
	"context"
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/MaximillianoNico/COVID-19-API/internals/controllers"
	"fmt"
	"github.com/joho/godotenv"
	
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connect()(*mongo.Database, error) {
	clientOptions := options.Client()
    clientOptions.ApplyURI(os.Getenv("MONGODB_URL"))
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Connect(context.TODO())
    if err != nil {
        return nil, err
	}
	
	fmt.Println("Connected to MLab MongoDB")

    return client.Database(os.Getenv("MONGODB_DB_NAME")), nil
}

func definedRoute (client *mongo.Database) (routes *mux.Router) {
	routes = mux.NewRouter()

	controllers.DbClient = client

	routes.HandleFunc("/api/:country/filters", controllers.GetDataCovid).Methods("GET")
	routes.HandleFunc("/api/symptoms/{country}", controllers.Telemedicine).Methods("GET")
	return
}

func main() {
	// Load Environment from file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connection to Database MongoDB
	client, err := connect() 

	if err != nil {
		log.Fatal("ERROR : ", err)
	}

	router := definedRoute(client)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), router))
}
