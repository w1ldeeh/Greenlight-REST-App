package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", sayHello)
	mux.HandleFunc("/bye", sayBye)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
