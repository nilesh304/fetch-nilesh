package api

import (
	"encoding/json"
	"fetch-project/app"
	"fetch-project/schema"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Process(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var reciept schema.Reciept
	err := json.NewDecoder(r.Body).Decode(&reciept)
	if err != nil {
		log.Printf("error %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "The receipt is invalid.",
		})
		return
	}

	if _, err := ValidateStruct(reciept); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(validationErrors) validation error helps identify which field has issue
		json.NewEncoder(w).Encode(map[string]string{
			"message": "The receipt is invalid.",
		})
		return
	}

	resp, err := app.Process(&reciept)
	if err != nil {
		log.Printf("error %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	// Extract the id from the URL path
	vars := mux.Vars(r)

	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	resp, err, status := app.GetPoint(id)
	if err != nil {
		log.Printf("error %s", err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
