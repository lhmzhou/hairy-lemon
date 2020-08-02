package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
