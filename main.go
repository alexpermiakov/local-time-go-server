package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	location, _ := time.LoadLocation("America/Vancouver")
	currentTime := time.Now().In(location).Format("Monday, January 2, 2006 at 3:04:05 PM MST")
	fmt.Fprintf(w, "Server Time: %s\n", currentTime)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ready")
}

func main() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ready", readyHandler)
	fmt.Println("Server starting on port 1599...")
	if err := http.ListenAndServe(":1599", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
