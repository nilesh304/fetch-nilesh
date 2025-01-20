package main

import (
	"fetch-project/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TimeStruct struct {
	CurrentTime string `json:"current_time"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/receipts/{id}/points", api.GetPoints).Methods("GET")
	r.HandleFunc("/receipts/process", api.Process).Methods("POST")

	// http.HandleFunc("/reciepts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
	// 	// Only allow Get requests
	// 	if r.Method != http.MethodGet {
	// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// 	vars := mux.Vars(r)
	// 	log.Printf("error %s", vars)

	// 	id, err := strconv.Atoi(vars["id"])
	// 	if err != nil {
	// 		log.Printf("error %s", err)
	// 		http.Error(w, "Error parsing id", http.StatusBadRequest)
	// 		return
	// 	}
	// 	resp, err := app.GetPoint(id)
	// 	if err != nil {
	// 		log.Printf("error %s", err)
	// 		http.Error(w, "Error saving reciept", http.StatusBadRequest)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusCreated)
	// 	json.NewEncoder(w).Encode(resp)
	// })
	// http.ListenAndServe(":8080", r)
	fmt.Println("Server is running at localhost: 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
