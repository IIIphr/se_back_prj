package Router

import (
	"se_back_prj/Controller"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/codes", Controller.FindCodes).Methods("POST")
	//router.HandleFunc("/api/selfs", Controller.GetAllWebsitesJSON).Methods("GET")
	router.HandleFunc("/api/login", Controller.CheckLogin).Methods("POST")
	router.HandleFunc("/api/signup", Controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/sell", Controller.CreateCoupon).Methods("POST")
	router.HandleFunc("/api/buy", Controller.DeleteOneCoupon).Methods("POST")
	router.HandleFunc("/api/admin", Controller.CreateAdmin).Methods("POST")
	router.HandleFunc("/api/university", Controller.CreateUniversity).Methods("POST")
	router.HandleFunc("/api/canteen", Controller.CreateCanteen).Methods("POST")
	router.HandleFunc("/api/report", Controller.CreateReport).Methods("POST")
	router.HandleFunc("/api/user", Controller.UpdateUser).Methods("POST")
	router.HandleFunc("/api/money", Controller.UserMoney).Methods("POST")
	router.HandleFunc("/api/update-money", Controller.UpdateMoney).Methods("POST")
	router.HandleFunc("/api/Feedback", Controller.GiveFeedback).Methods("POST")
	return router
}
