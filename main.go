package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jrexiti/backend-go/pkg/server"
)

func main() {

	address := ":8080"
	mux := http.NewServeMux()
	svr := server.New()

	mux.HandleFunc("/", svr.HandleIndex)
	mux.HandleFunc("/user", svr.HandleUsers)

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Start server: %v", address)
	log.Fatal(s.ListenAndServe())
}
