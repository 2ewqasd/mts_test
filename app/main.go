package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": 200}`))
	case "POST":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": 200}`))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "method not allowed"}`))
	}
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":80", nil))
}
