package main

import (
  "os"
  "log"
  "fmt"
  "io/ioutil"
  "net/http"
  "github.com/gorilla/mux"
)

const (
	serverport = ":80"
  api_name = "API1"
)

var (
  // api to call
  api_to_call = os.Getenv("API_TO_CALL")
)

//--------------------------------
func handleRequests() {

		// set up routes
		router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homePage).Methods("GET")
    router.HandleFunc("/child1", child1).Methods("GET")
    router.HandleFunc("/callApi/home", callApiHome).Methods("GET")

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

//--------------------------------
func callApiHome(w http.ResponseWriter, r *http.Request) {

    log.Println("INFO - callApiHome(): Endpoing hit - callapi.  Calling API: " + api_to_call)

    response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
    if err != nil {
      log.Println("ERR - callApiHome(): Error Calling API: " + api_to_call + " :" + err.Error())
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
      log.Println("ERR - callApiHome(): Error reading response from API: " + api_to_call + " :" + err.Error())
    }

    //fmt.Println(string(responseData))
    //fmt.Fprintf(w, responseData)
    err := r.Write(response)
    if err != nil {
      log.Println("ERR - callApiHome(): Error writing response: " + err.Error())
      fmt.Fprintf(w, "ERR - Could not write here the response received from the called API")
    }

}

func main() {
  handleRequests()
}
