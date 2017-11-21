package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Post() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			log.Println(time.Now(), " : POST Request recieved.")

			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				fmt.Println("Error occured while reading request body stream :", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Printf("\n\n%v\n\n", string(body))
			var v interface{}
			err = json.Unmarshal(body, &v)

			//fmt.Println(v)

			if err != nil {
				fmt.Println("Error occured while unmarshaling/Parsing JSON request body : ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello from POST handler!!"))
		}

	})
}
