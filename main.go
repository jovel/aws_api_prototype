package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Messages!\n"))
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/messages", GetMessages)
	r.HandleFunc("/v1/healthcheck", Healthcheck)


	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":50800", r))
}