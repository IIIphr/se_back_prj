package Router

import (
	"se_back_prj/Controller"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/codes/{self}", Controller.GetAllWebsitesJSON).Methods("GET")
	router.HandleFunc("/api/selfs", Controller.GetAllWebsitesJSON).Methods("GET")
	router.HandleFunc("/api/login", Controller.CreateWebsite).Methods("POST")
	router.HandleFunc("/api/signup", Controller.CreateWebsite).Methods("POST")
	router.HandleFunc("/api/sell/{self}+{code}", Controller.MarkAsChecked).Methods("PUT")
	router.HandleFunc("/api/codes/{id}", Controller.DeleteOneWebsite).Methods("DELETE")
	return router
}
