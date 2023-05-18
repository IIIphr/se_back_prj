package Controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"se_back_prj/Model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://avatar:kami1380@cluster0.ibwwj5y.mongodb.net/test"
const dbName = "dblab"
const colName = "id"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection is ready")
}

func insertOneQuery(website Model.WebSite) {
	inserted, err := collection.InsertOne(context.Background(), website)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted website with id ", inserted.InsertedID, " into db")
}

func updateWebSite(website string) {
	id, _ := primitive.ObjectIDFromHex(website)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"checked": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated ", result.ModifiedCount)
}
func deleteOneWebSite(website string) {
	id, _ := primitive.ObjectIDFromHex(website)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete count was ", deleteCount.DeletedCount)
}
func deleteAll() int64 {

	deleteNum, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted all websites ", deleteNum.DeletedCount)
	return deleteNum.DeletedCount
}

func getAllWebsites() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var websites []primitive.M
	for cur.Next(context.Background()) {
		var website bson.M
		err := cur.Decode(&website)
		if err != nil {
			log.Fatal(err)
		}
		websites = append(websites, website)
	}
	defer cur.Close(context.Background())
	return websites
}

func GetAllWebsitesJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allWebsites := getAllWebsites()
	json.NewEncoder(w).Encode(allWebsites)
}

func CreateWebsite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var website Model.WebSite
	_ = json.NewDecoder(r.Body).Decode(&website)
	insertOneQuery(website)
	json.NewEncoder(w).Encode(website)
}

func MarkAsChecked(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	updateWebSite(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneWebsite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneWebSite(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllWebsite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAll()
	json.NewEncoder(w).Encode(count)
}
