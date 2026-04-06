package db

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://singharkirath1511_db_user:Harkirat%40123@cluster.e6wscec.mongodb.net/?appName=Cluster"
const dbName = "netflix"
const colName = "watchlist"

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("Noooo!!! \n err : ", err)
	}
	fmt.Println("Mongodb connection success")

	Collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}
