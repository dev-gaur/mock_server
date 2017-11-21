package handlers

import (
	"log"
	"net/http"
	"time"
)

func Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(time.Now(), " : GET Request recieved.")

		key := r.URL.Query().Get("key")

		defer r.Body.Close()

		if key != "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from Get handler!! Key :"))
	})
}
