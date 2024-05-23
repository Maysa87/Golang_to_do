package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Maysa87/Goland_to_do/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.collention

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loadind the .env file")
	}
}
func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApllyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to MongoDB!")
	client.Database(dbName).collection(collectionName)
	fmt.Println("collection instance created")
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-type")
	var Task models.ToDo
	json.NewDecoder(r.Body).Decode(&Task)
	insertOneTask(Task)
	json.NewEncoder(w).Encode(Task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-type")
	params := mux.Vars(r)
	TaskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-type")
	params := mux.Vars(r)
	UndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
}

func getAllTask() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cursor.Next(context.Background()) {
		var result bson.M
		e := cursor.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cursor.err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.Background())
	return results
}

func taskComplete(Task string) {
	id, _ := primitive.ObjectIDFromHeaders
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), update, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)
}

func insertOneTask(Task models.ToDo) {
	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single record", insertResult.InsertedID)
}

func undoTask(Task String) {
	id, _ := primitive.objectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), update, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count:", result.ModifiedCount)
}

func deleteOneTask(Task String) {
	id, _ := primitive.objectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted document:", result.DeleteCount)
}

func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), "", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted document:", result.DeleteCount)
	return d.DeleteCount
}
