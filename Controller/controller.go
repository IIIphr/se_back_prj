package Controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"se_back_prj/Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var adminCollection *mongo.Collection
var couponCollection *mongo.Collection
var canteenCollection *mongo.Collection
var reportCollection *mongo.Collection
var universityCollection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	userCollection = client.Database("food").Collection("users")
	adminCollection = client.Database("food").Collection("admins")
	couponCollection = client.Database("food").Collection("coupons")
	canteenCollection := client.Database("food").Collection("canteens")
	reportCollection := client.Database("food").Collection("reports")
	universityCollection := client.Database("food").Collection("universities")
	_ = canteenCollection
	_ = reportCollection
	_ = universityCollection
	fmt.Println("collection is ready")
	fmt.Println(universityCollection.Name())
}
func insertNewUser(user Model.User) {
	inserted, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted user with id ", inserted.InsertedID, " into db")
}
func insertNewCoupon(coupon Model.Coupon) {
	inserted, err := couponCollection.InsertOne(context.Background(), coupon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted coupon with id ", inserted.InsertedID, " into db")
}
func deleteId(delete Model.DeletingCoupon) Model.CurStatus {
	id, _ := primitive.ObjectIDFromHex(string(delete.ID))
	filter := bson.M{"_idcoupon": id}
	var result Model.Coupon
	err := couponCollection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// User not found
		var stat Model.CurStatus
		stat.Stat = "Coupon Not Found"
		return stat
	} else if err != nil {
		log.Fatal(err)
	}
	result.Owner.CurrentMoney += result.Price
	delete.Buyer.CurrentMoney -= result.Price
	fmt.Println("delete count was ", result)
	var stat Model.CurStatus
	stat.Stat = "FOUND"
	return stat
}
func insertNewCAdmin(admin Model.Admin) {
	inserted, err := adminCollection.InsertOne(context.Background(), admin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted admin with id ", inserted.InsertedID, " into db")
}
func insertNewUniversity(university Model.University) {
	inserted, err := universityCollection.InsertOne(context.Background(), university)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted university with id ", inserted.InsertedID, " into db")
}
func insertNewCanteen(canteen Model.Canteen) {
	inserted, err := canteenCollection.InsertOne(context.Background(), canteen)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted Canteen with id ", inserted.InsertedID, " into db")
}
func insertNewReport(report Model.Report) {
	inserted, err := reportCollection.InsertOne(context.Background(), report)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted report with id ", inserted.InsertedID, " into db")
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
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var admin Model.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)
	insertNewCAdmin(admin)
}
func CreateCanteen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var canteen Model.Canteen
	_ = json.NewDecoder(r.Body).Decode(&canteen)
	insertNewCanteen(canteen)
}
func CreateUniversity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var university Model.University
	_ = json.NewDecoder(r.Body).Decode(&university)
	insertNewUniversity(university)
}
func CreateReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var report Model.Report
	_ = json.NewDecoder(r.Body).Decode(&report)
	insertNewReport(report)
}

func DeleteOneCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	var delete Model.DeletingCoupon
	err := json.NewDecoder(r.Body).Decode(&delete)
	if err != nil {
		var stat Model.CurStatus
		stat.Stat = "INVALID"
		json.NewEncoder(w).Encode(stat)
		return
	}
	json.NewEncoder(w).Encode(deleteId(delete))
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var stat Model.CurStatus
		stat.Stat = "INVALID"
		json.NewEncoder(w).Encode(stat)
		return
	}
	filter := bson.M{"studentid": user.StudentId, "password": user.Password}
	var result Model.User
	err = userCollection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// User not found
		var stat Model.CurStatus
		stat.Stat = "NOT OK"
		json.NewEncoder(w).Encode(stat)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
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
				"canteen": self,
			},
		},
		{
			"$sort": bson.M{
				"price": 1,
			},
		},
	}
	cursor, err := couponCollection.Aggregate(context.Background(), pipeline)
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

//TODO login user model sent to front, change current money of user.
