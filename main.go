package main

import (
	"github.com/gorilla/mux"
	"hairy-lemon/producer"
	"log"
	"net/http"
)

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/score/{id}", producer.GetById)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
