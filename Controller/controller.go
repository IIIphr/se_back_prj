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

const connectionString = "mongodb://localhost:27017"
const dbName = "samad2"
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

func insertNewUser(user Model.User) {
	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted user with id ", inserted.InsertedID, " into db")
}
func insertNewCoupon(coupon Model.Coupon) {
	inserted, err := collection.InsertOne(context.Background(), coupon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted coupon with id ", inserted.InsertedID, " into db")
}
func deleteId(val string) Model.CurStatus {
	id, _ := primitive.ObjectIDFromHex(val)
	filter := bson.M{"_idcoupon": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			var stat Model.CurStatus
			stat.Stat = "Coupon Not Found"
			return stat
		}
		log.Fatal(err)
	}
	fmt.Println("delete count was ", deleteCount.DeletedCount)
	var stat Model.CurStatus
	stat.Stat = "FOUND"
	return stat
}

func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var coupon Model.Coupon
	_ = json.NewDecoder(r.Body).Decode(&coupon)
	insertNewCoupon(coupon)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertNewUser(user)
}
func DeleteOneCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(deleteId(params["id"]))
}
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	filter := bson.M{"studentid": user.StudentId, "password": user.Password}
	var result Model.User
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// User not found
		var stat Model.CurStatus
		stat.Stat = "NOT OK"
		json.NewEncoder(w).Encode(stat)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	var stat Model.CurStatus
	stat.Stat = "OK"
	json.NewEncoder(w).Encode(stat)
}
func FindCodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var self string
	err := json.NewDecoder(r.Body).Decode(&self)
	if err != nil {
		var stat Model.CurStatus
		stat.Stat = "INVALID"
		json.NewEncoder(w).Encode(stat)
		return
	}
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"self": self,
			},
		},
		{
			"$sort": bson.M{
				"price": 1,
			},
		},
	}
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the results
	var results []Model.Coupon
	for cursor.Next(context.Background()) {
		var coupon Model.Coupon
		err := cursor.Decode(&coupon)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, coupon)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	var response []Model.Coupon
	response = append(response, results[0])
	response = append(response, results[1])
	response = append(response, results[2])
	response = append(response, results[3])
	response = append(response, results[4])
	json.NewEncoder(w).Encode(response)
}
