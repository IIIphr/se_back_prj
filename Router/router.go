package Router

import (
	"se_back_prj/Controller"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/codes", Controller.FindCodes).Methods("POST") //model = canteen
	//router.HandleFunc("/api/selfs", Controller.GetAllWebsitesJSON).Methods("GET")
	router.HandleFunc("/api/login", Controller.CheckLogin).Methods("POST")              //model = user
	router.HandleFunc("/api/signup", Controller.CreateUser).Methods("POST")             //model = user
	router.HandleFunc("/api/sell", Controller.CreateCoupon).Methods("POST")             //model = coupon
	router.HandleFunc("/api/buy", Controller.DeleteOneCoupon).Methods("POST")           //model =deletingcoupon
	router.HandleFunc("/api/admin", Controller.CreateAdmin).Methods("POST")             //model =admin
	router.HandleFunc("/api/university", Controller.CreateUniversity).Methods("POST")   //model =university
	router.HandleFunc("/api/canteen", Controller.CreateCanteen).Methods("POST")         //model =cantenn
	router.HandleFunc("/api/report", Controller.CreateReport).Methods("POST")           //model =report
	router.HandleFunc("/api/user", Controller.UpdateUser).Methods("POST")               //model =user
	router.HandleFunc("/api/money", Controller.UserMoney).Methods("POST")               //model = user (returns userMoney)
	router.HandleFunc("/api/update-money", Controller.UpdateMoney).Methods("POST")      //model =UpdateMoney
	router.HandleFunc("/api/feedback", Controller.GiveFeedback).Methods("POST")         //model =Feedback
	router.HandleFunc("/api/transfer-money", Controller.TransferMoney).Methods("POST")  //model =MoneyTransfer
	router.HandleFunc("/api/history", Controller.History).Methods("POST")               //model =user (returns UserHistory)
	router.HandleFunc("/api/get-universities", Controller.Universities).Methods("POST") //model =nothing (returns array of University)
	router.HandleFunc("/api/get-canteens", Controller.Canteens).Methods("POST")         //model =University (returns array of Canteen)
	return router
}
