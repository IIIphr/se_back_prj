package Controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"se_back_prj/Manager"
	"se_back_prj/Model"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var adminCollection *mongo.Collection
var couponCollection *mongo.Collection
var canteenCollection *mongo.Collection
var reportCollection *mongo.Collection
var universityCollection *mongo.Collection
var CounterCollection *mongo.Collection

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
	canteenCollection = client.Database("food").Collection("canteens")
	reportCollection = client.Database("food").Collection("reports")
	universityCollection = client.Database("food").Collection("universities")
	CounterCollection = client.Database("food").Collection("counter")
	fmt.Println("collection is ready")
	fmt.Println(universityCollection)
}
func insertNewUser(user Model.User) {
	inserted, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted user with id ", inserted.InsertedID, " into db")
}
func updateName(user Model.User) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"firstname": user.FirstName}}
	result, err := userCollection.UpdateOne(context.Background(), filter, update)
	update2 := bson.M{"$set": bson.M{"firstname": user.FirstName}}
	result2, err2 := userCollection.UpdateOne(context.Background(), filter, update2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated ", result.ModifiedCount)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("updated ", result2.ModifiedCount)
}
func updateUserPassword(user Model.User) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"password": user.Password}}
	result, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated ", result.ModifiedCount)
}
func insertNewCoupon(coupon Model.Coupon) {
	coupon.ID = Manager.GetNewCouponId()
	inserted, err := couponCollection.InsertOne(context.Background(), coupon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted coupon with id ", inserted.InsertedID, " into db")
}
func updateUserMoney(user Model.User, money int) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"currentmoney": money}}
	result, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated ", result.ModifiedCount)
}
func deleteId(ID int64, buyersid string, buyeruid string) Model.CurStatus {
	id := ID
	filter := bson.M{"_id": id}
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
	var res Model.User
	filter2 := bson.M{"studentid": buyersid, "universityid": buyeruid}
	userCollection.FindOne(context.Background(), filter2).Decode(&res)
	//res.CurrentMoney -= result.Price
	updateUserMoney(res, res.CurrentMoney-result.Price)
	var res2 Model.User
	filter3 := bson.M{"studentid": result.StudentId, "universityid": result.University}
	userCollection.FindOne(context.Background(), filter3).Decode(&res2)
	//res2.CurrentMoney += result.Price
	updateUserMoney(res2, res2.CurrentMoney+result.Price)

	filter4 := bson.M{"studentid": buyersid, "universityid": buyeruid}
	update := bson.M{
		"$push": bson.M{
			"coupons": result,
		},
	}
	updateResult, err := userCollection.UpdateOne(context.TODO(), filter4, update)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Matched %v document(s) and updated %v document(s).\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	deleteResult, err := couponCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete was ", result, deleteResult)
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
func getUserMoney(user Model.User) Model.UserMoney {
	var result Model.UserMoney

	filter := bson.M{"studentid": user.StudentId, "password": user.Password}
	var result2 Model.User
	_ = userCollection.FindOne(context.Background(), filter).Decode(&result2)
	result.CurrentMoney = result2.CurrentMoney
	return result
}
func UpdateMoney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.UpdateMoney
	_ = json.NewDecoder(r.Body).Decode(&user)
	var money = getUserMoney(user.User).CurrentMoney
	updateUserMoney(user.User, user.Money+money)
}
func UserMoney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(getUserMoney(user))
}
func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var coupon Model.Coupon
	_ = json.NewDecoder(r.Body).Decode(&coupon)
	insertNewCoupon(coupon)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertNewUser(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user Model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	updateName(user)
	updateUserPassword(user)
}
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var admin Model.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)
	insertNewCAdmin(admin)
}
func CreateCanteen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var canteen Model.Canteen
	_ = json.NewDecoder(r.Body).Decode(&canteen)
	insertNewCanteen(canteen)
}
func CreateUniversity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var university Model.University
	_ = json.NewDecoder(r.Body).Decode(&university)
	insertNewUniversity(university)
}
func CreateReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var report Model.Report
	_ = json.NewDecoder(r.Body).Decode(&report)
	insertNewReport(report)
}

func DeleteOneCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var delete Model.DeletingCoupon
	err := json.NewDecoder(r.Body).Decode(&delete)
	if err != nil {
		var stat Model.CurStatus
		stat.Stat = "INVALID"
		json.NewEncoder(w).Encode(stat)
		return
	}
	json.NewEncoder(w).Encode(deleteId(delete.ID, delete.BuyerStudentID, delete.BuyerUniversityID))
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var self Model.Canteen
	err := json.NewDecoder(r.Body).Decode(&self)
	if err != nil {
		var stat Model.CurStatus
		stat.Stat = "INVALID"
		json.NewEncoder(w).Encode(stat)
		return
	}
	filter := bson.M{"canteenid": self.ID}
	cur, err := couponCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var coupons []Model.Coupon
	for cur.Next(context.Background()) {
		var coupon Model.Coupon
		err := cur.Decode(&coupon)
		if err != nil {
			log.Fatal(err)
		}
		coupons = append(coupons, coupon)
	}
	sort.Sort(ByPrice(coupons))
	var response []Model.Coupon
	response = append(response, coupons[0])
	response = append(response, coupons[1])
	response = append(response, coupons[2])
	response = append(response, coupons[3])
	response = append(response, coupons[4])
	json.NewEncoder(w).Encode(response)
}

type ByPrice []Model.Coupon

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }
