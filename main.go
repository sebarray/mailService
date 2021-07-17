package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sebarray/mailService/met"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", met.SendMail).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:2000", router))

}
