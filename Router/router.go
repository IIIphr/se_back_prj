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
	router.HandleFunc("/api/sell", Controller.CreateCoupon).Methods("PUT")
	router.HandleFunc("/api/codes", Controller.DeleteOneCoupon).Methods("DELETE")
	return router
}
