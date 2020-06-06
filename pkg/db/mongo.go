// package database

// import (
// 	"context"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var (
// 	connectTimeout		= 5
// 	ctxBg 				= context.Background()
// )

// func Connect() (*mongo.Database, error) {
// 	clientOptions := options.Client()
// 	clientOptions.ApplyURI("mongodb://admin:admin123@ds145299.mlab.com:45299/mext")
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.Database("mext"), ctx, nil
// }

// Define the package interacting with the database
package db

import (
	"os"
	"time"
	"fmt"
	"gopkg.in/mgo.v2"
)

// DBConnection defines the connection structure
type DBConnection struct {
	session *mgo.Session
}

// NewConnection handles connecting to a mongo database
func NewConnection(host string, dbName string) (conn *DBConnection) {
	fmt.Println("enter main - connecting to mongo")
	
	info := &mgo.DialInfo{
		// Address if its a local db then the value host=localhost
		Addrs: []string{host},
		// Timeout when a failure to connect to db
		Timeout: 60 * time.Second,
		// Database name
		Database: dbName,
		// Database credentials if your db is protected
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}
	return conn
}

// Use handles connect to a certain collection
func (conn *DBConnection) Use(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *DBConnection) Close() {
	// This closes the connection
	conn.session.Close()
	return
}