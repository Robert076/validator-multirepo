package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Robert076/validator-multirepo/internal/data"
	"github.com/Robert076/validator-multirepo/internal/validator"
)

func main() {
	var serviceName string = "VALIDATOR"

	http.HandleFunc("/validator", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed on this endpoint.", http.StatusMethodNotAllowed)
			log.Printf("%s: Only POST method is allowed on this endpoint.", serviceName)
			return
		}

		var body data.ExpectedBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid JSON format.", http.StatusBadRequest)
			log.Printf("%s: Expected decoding the request body.", serviceName)
			return
		}

		if valid, err := validator.IsNameValid(body.Name); err != nil {
			http.Error(w, "Error occured when checking if name is valid or not. ", http.StatusInternalServerError)
			log.Printf("%s: Error occured when checking if name is valid or not. %s", serviceName, err)
			return
		} else if !valid {
			http.Error(w, "Name is invalid. Make sure it contains no numbers.", http.StatusBadRequest)
			log.Printf("%s: %s is not a valid name", serviceName, body.Name)
			return
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("%s: Failed to start server.", serviceName)
	}
}
