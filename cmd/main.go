package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Robert076/validator-multirepo/internal/validator"
)

func main() {
	var serviceName string = "VALIDATOR"

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("%s: Failed to start server.", serviceName)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed on this endpoint.", http.StatusMethodNotAllowed)
			log.Printf("%s: Only POST method is allowed on this endpoint.", serviceName)
			return
		}

		var name string

		if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
			http.Error(w, "Invalid JSON format.", http.StatusBadRequest)
			log.Printf("%s: Expected the property `name` in the JSON, but could not find it.", serviceName)
			return
		}

		if valid, err := validator.IsNameValid(name); err != nil {
			http.Error(w, "Error occured when checking if name is valid or not.", http.StatusInternalServerError)
			log.Printf("%s: Error occured when checking if name is valid or not.", serviceName)
		}
	})
}
