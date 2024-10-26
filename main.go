package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Middleware untuk mencetak log request
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}

// Handler untuk (/jakarta) endpoint
func jakartaHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the jakarta page!")
}

// Handler untuk (/bali) endpoint
func baliHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the bali page!")
}

// Handler untuk (/papua) endpoint
func papuaHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the papua page!")
}

func main(){
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/", fileServer).Methods("GET")
	r.HandleFunc("/jakarta", jakartaHandler).Methods("GET")
	r.HandleFunc("/bali", baliHandler).Methods("GET")
	r.HandleFunc("/papua", papuaHandler).Methods("GET")

	fmt.Println("-------------------------------------------")
	fmt.Println("| Server running on http://localhost:8090 |")
	fmt.Println("-------------------------------------------")
	fmt.Println("Endpoint 1 -> http://localhost:8090/jakarta")
	fmt.Println("Endpoint 2 -> http://localhost:8090/bali")
	fmt.Println("Endpoint 3 -> http://localhost:8090/papua")
	fmt.Println("-------------------------------------------")
	fmt.Println("> Log request")
	fmt.Print("")

	if err := http.ListenAndServe(":8090", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}