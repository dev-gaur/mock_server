package main

import (
	"log"
	"hdd/mock_server/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/get", handlers.Get())
	mux.Handle("/post", handlers.Post())

	log.Fatal(http.ListenAndServe(":8888", mux))
}
