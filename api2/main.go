package main

import (
  "log"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

const (
	serverport = ":80"
  api_name = "API2"
)



//--------------------------------
func handleRequests() {

		// set up routes
		router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homePage).Methods("GET")
    router.HandleFunc("/child1", child1).Methods("GET")


		// start the server
    log.Fatal(http.ListenAndServe(serverport, router))
}

//--------------------------------
func homePage(w http.ResponseWriter, r *http.Request) {
    log.Println("INFO - homePage(): Endpoing hit - homepage")
    fmt.Fprintf(w, "Welcome to " + api_name) // writes to httpwriter
}

//--------------------------------
func child1(w http.ResponseWriter, r *http.Request) {
    log.Println("INFO - child1(): Endpoing hit - child1")
    fmt.Fprintf(w, "Welcome to CHILD1 of " + api_name) // writes to httpwriter
}

func main() {
  handleRequests()
}
