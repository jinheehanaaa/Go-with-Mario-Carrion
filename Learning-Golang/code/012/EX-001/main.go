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
// 003. Functional Options
// Introduce Option Type that declares function
// We return this Option

type Option func(*http.Server)

func WithAddr(addr string) func(*http.Server) {
	return func(s *http.Server) {
		s.Addr = addr
	}
}

func WithReadTimeout(t time.Duration) func(*http.Server) {
	return func(s *http.Server) {
		s.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) func(*http.Server) {
	return func(s *http.Server) {
		s.WriteTimeout = t
	}
}

func NewServer(opts ...Option) http.Server {
	s := http.Server{
		//Addr:         opt.Address,
		//ReadTimeout:  opt.ReadTimeout,
		//WriteTimeout: opt.WriteTimeout,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	s := NewServer(
		WithAddr(":9191"),
		WithReadTimeout(100*time.Millisecond),
		WithWriteTimeout(100*time.Millisecond),
	)
	s.Handler = mux

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and server", err)
	}
}
