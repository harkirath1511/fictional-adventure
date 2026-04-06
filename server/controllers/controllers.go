package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harkirath1511/mongo-api/db"
	"github.com/harkirath1511/mongo-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var collection = db.Collection

//mongo helper functions

// insert 1
func insertOneMovie(movie models.Netflix) {
	res, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("ERR : ", err)
	}

	fmt.Println("Inserted one movie in db with id : ", res.InsertedID)
}

// update 1
func updateOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal("The given id id not valid or some err : ", err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"iswatched": true}}

	res, _ := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("Modified cnt = ", res.ModifiedCount)
}

// delete 1
func deleteOne(movieId string) {
	id, _ := bson.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("Given id not valid or some err : ", err)
	}

	fmt.Println("Deleted cnt : ", res.DeletedCount)
}

// delete all
func deleteAll() int64 {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal("err deleting : ", err)
	}
	fmt.Println("Number of movies deleted : ", res.DeletedCount)
	return res.DeletedCount
}

func getAll() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal("Some err: ", err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, primitive.M(movie))
	}
	defer cursor.Close(context.Background())
	return movies
}

//actual controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAll()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	deleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cnt := deleteAll()
	json.NewEncoder(w).Encode(cnt)
}
