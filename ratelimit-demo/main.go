package main

import (
	"log"
	"net/http"
	"time"
)

func temp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)
	log.Println("Listening at :8080...")
	http.ListenAndServe(":8080", ratelimit(mux))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	w.Write([]byte("OK"))
}
