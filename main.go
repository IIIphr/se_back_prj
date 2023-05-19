package main

import (
	"fmt"
	"log"
	"net/http"

	//"se_back_prj/Controller"
	"se_back_prj/Router"
	//"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("mongo api")
	r := Router.Route()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at 4000")

	//admin := bson.D{{"adminid", "1"}, {"username", "rootadmin"}, {"password", "root"}, {"name", "Arash"}}
	//user := bson.D{{"studentid", "98243001"}, {"universityid", "1"}, {"FirstName", "Sepehr"}, {"LastName", "Ebrahimi"}, {"Password", "123"}, {"CurrentMoney", 500}}
	//report := bson.D{{"reportedcoupon", "1"}, {"reporter", "98243001"}, {"reportee", "98243002"}}
	//university := bson.D{{"universityid", "1"}}
	//coupon := bson.D{{"code", "54321"}, {"price", 2}, {"canteen", "2"}, {"foodname", "قورمه سبزی"}, {"Owner", user}}
	//canteen := bson.D{{"canteenid", "2"}, {"universityid", "1"}}
}
