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
// 002. Change with pattern
// Default config value from outisde of main
// Use ServerOpts for Data Transfer Object (DTO) config

type ServerOpts struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Take default value from ServerOpts
func NewServerOpts() ServerOpts {
	return ServerOpts{
		Address:      ":8080",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}

func NewServer(opt ServerOpts) http.Server {
	s := http.Server{
		Addr:         opt.Address,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
	}

	return s
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	s := NewServer(NewServerOpts())

	s.Handler = mux

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and server", err)
	}
}
