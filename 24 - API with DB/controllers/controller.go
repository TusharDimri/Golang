package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/TusharDimri/GolangApiWithDB/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "Netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// Connect with MongoDB
func init() {
	// Client Options:
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB:
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	// Collection Instance
	fmt.Println("Collection reference is ready")
}

//  Mongodb Helpers - File

// Insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with id:", inserted.InsertedID)
}

// Update 1 record
func updateOneMovie(movieId string) {
	_id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// Delete 1 record
func deleteOneMovie(movieId string) {
	_id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": _id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie go deleted with delete count: ", deleteCount)
}

// Delete all records
func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all movies from db
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

		defer cur.Close(context.Background())
	}
	return movies
}

//  Actual Controllers -file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
