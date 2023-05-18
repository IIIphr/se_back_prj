package Router

import (
	"se_back_prj/Controller"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/websites", Controller.GetAllWebsitesJSON).Methods("GET")
	router.HandleFunc("/api/website", Controller.CreateWebsite).Methods("POST")
	router.HandleFunc("/api/website/{id}", Controller.MarkAsChecked).Methods("PUT")
	router.HandleFunc("/api/website/{id}", Controller.DeleteOneWebsite).Methods("DELETE")
	router.HandleFunc("/api/websites", Controller.DeleteAllWebsite).Methods("DELETE")
	return router
}
