package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 001. HTTP Server Implementation
// curl -X GET http://localhost:8080/
// curl -X GET http://127.0.0.1:8080/

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	s := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and server", err)
	}
}
