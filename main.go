package main

import (
	"fmt"
	"log"
	"net/http"
	"se_back_prj/Router"
)

func main() {
	fmt.Println("mongo pi")
	r := Router.Route()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at 4000")
}
