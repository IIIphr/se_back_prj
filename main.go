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
	usersCollection := client.Database("food").Collection("users")
	adminCollection := client.Database("food").Collection("admins")
	couponCollection := client.Database("food").Collection("coupons")
	canteenCollection := client.Database("food").Collection("canteens")
	reportCollection := client.Database("food").Collection("reports")
	universityCollection := client.Database("food").Collection("universities")

}
