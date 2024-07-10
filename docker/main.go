package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleGreetingEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Zomato!")
	log.Printf("New request on %s", r.URL)
}

func startHelloApp() {
	log.Println("Application started successfully")
	http.HandleFunc("/hello", handleGreetingEndpoint)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatalf("Failed to start the app. Error: %s", err)
	}
	log.Fatal()
}

func main() {
	startHelloApp()
}
