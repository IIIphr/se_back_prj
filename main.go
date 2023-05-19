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
	admin:= bson.D{{"adminID","1"},{"username","rootadmin"},{"password","root"},{"Name","Arash"}}
	result, err := adminCollection.InsertOne(context.TODO(), admin)
	if err != nil {
        panic(err)
	}
	user:= bson.D{{"StudentID","98243001"},{"UniversityID","1"},{"FirstName","Sepehr"},{"LastName","Ebrahimi"},{"password","123"},{"credit",500}}
	result, err:= userCollection.InsertOne(context.TODO(), user)
	if err != nil{
		panic(err)
	}
	report:= bson.D{{"reportID","1"},{"couponID","1",}{"StudentID","98243001"}}
	result, err:= reportCollection.InsertOne(context.TODO(), report)
	if err != nil{
		panic(err)
	}
	university:=bson.D{{"UniversityID","1"}}
	result, err:= universityCollection.InsertOne(context.TODO(), university)
	if err != nil{
		panic(err)
	}
	coupon:=bson.D{{"CouponID","1"},{"Code","54321"},{"Price",2},{"CanteenID","2"},{"FoodName","قورمه سبزی"}}
	result, err:= couponCollection.InsertOne(context.TODO(), coupon)
	if err != nil{
		panic(err)
	}
	canteen:= bson.D{{"CanteenID","2"}}
	result, err:= canteenCollection.InsertOne(context.TODO(), canteen)
	if err != nil{
		panic(err)
	}

}
