package main

import (
	"log"
	"net/http"
	"servicemail/met"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", met.SendMail).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:2000", router))

}
