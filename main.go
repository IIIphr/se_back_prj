package main

import (
	"fmt"
	"log"
	"net/http"
	"se_back_prj/Router"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
        if err != nil {
                panic(err)
        }
	fmt.Println("mongo pi")
	r := Router.Route()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at 4000")
	userCollection := client.Database("food").Collection("users")
	adminCollection := client.Database("food").Collection("admins")
	couponCollection := client.Database("food").Collection("coupons")
	canteenCollection := client.Database("food").Collection("canteens")
	reportCollection := client.Database("food").Collection("reports")
	universityCollection := client.Database("food").Collection("universities")
	admin:= bson.D{{"adminid","1"},{"username","rootadmin"},{"password","root"},{"name","Arash"}}
	result, err := adminCollection.InsertOne(context.TODO(), admin)
	if err != nil {
        panic(err)
	}
	user:= bson.D{{"studentid","98243001"},{"universityid","1"},{"FirstName","Sepehr"},{"LastName","Ebrahimi"},{"Password","123"},{"CurrentMoney",500}}
	result, err:= userCollection.InsertOne(context.TODO(), user)
	if err != nil{
		panic(err)
	}
	report:= bson.D{{"reportedcoupon","1",},{"reporter","98243001"},{"reportee","98243002"}}
	result, err:= reportCollection.InsertOne(context.TODO(), report)
	if err != nil{
		panic(err)
	}
	university:=bson.D{{"universityid","1"}}
	result, err:= universityCollection.InsertOne(context.TODO(), university)
	if err != nil{
		panic(err)
	}
	coupon:=bson.D{{"code","54321"},{"price",2},{"canteen","2"},{"foodname","قورمه سبزی"},{"Owner",user}}
	result, err:= couponCollection.InsertOne(context.TODO(), coupon)
	if err != nil{
		panic(err)
	}
	canteen:= bson.D{{"canteenid","2"},{"universityid","1"}}
	result, err:= canteenCollection.InsertOne(context.TODO(), canteen)
	if err != nil{
		panic(err)
	}

}
