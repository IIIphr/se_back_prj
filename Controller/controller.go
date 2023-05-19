package Controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"se_back_prj/Model"
	"sort"
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
func insertNewCoupon(ocupon Model.Coupon) {
	inserted, err := collection.InsertOne(context.Background(), coupon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted coupon with id ", inserted.InsertedID, " into db")
}
func deleteId(val string) {
	id, _ := primitive.ObjectIDFromHex(val)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete count was ", deleteCount.DeletedCount)
}

func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var coupon Model.Coupon
	_ = json.NewDecoder(r.Body).Decode(&coupon)
	insertNewCoupon(coupon)
	json.NewEncoder(w).Encode(coupon)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertNewUser(user)
	json.NewEncoder(w).Encode(user)
}
func DeleteOneCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteId(params["id"])
	json.NewEncoder(w).Encode(params["id"])
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
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Fatal(err)
	}

	// User found
	w.WriteHeader(http.StatusOK)
}
func FindCodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var self string
	err := json.NewDecoder(r.Body).Decode(&self)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	filter := bson.M{"self": self}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(cur, func(i, j int) bool {
		return cur[i].Price < cur[j].Price
	})
	var selfs []Model.Self
	for cur.Next(context.Background()) {
		var self Model.Self
		err := cur.Decode(&self)
		if err != nil {
			log.Fatal(err)
		}
		selfs = append(selfs, self)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the users as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selfs)
}
