package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler buat (/) endpoint
func homeHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

// Handler buat (/jakarta) endpoint
func jakartaHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the jakarta page!")
}

// Handler buat (/bali) endpoint
func baliHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the bali page!")
}

// Handler buat (/papua) endpoint
func papuaHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the papua page!")
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/jakarta", jakartaHandler).Methods("GET")
	r.HandleFunc("/bali", baliHandler).Methods("GET")
	r.HandleFunc("/papua", papuaHandler).Methods("GET")

	fmt.Println("-------------------------------------------")
	fmt.Println("| Server running on http://localhost:8090 |")
	fmt.Println("-------------------------------------------")
	fmt.Println("Endpoint 1 -> http://localhost:8090/jakarta")
	fmt.Println("Endpoint 2 -> http://localhost:8090/bali")
	fmt.Println("Endpoint 3 -> http://localhost:8090/papua")
	fmt.Print("> ")

	http.ListenAndServe(":8090", r)
}