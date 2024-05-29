package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Beadko/mywinebook/data"
	"github.com/Beadko/mywinebook/internal/wine"
	"github.com/gorilla/mux"
)

// AddRouterEndpoints add the actual endpoints for api
func AddRouterEndpoints(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/wine", getWine).Methods("GET")
	r.HandleFunc("/api/wine", addWine).Methods("POST")
	r.HandleFunc("/api/wine/{WINE_ID}", deleteWine).Methods("DELETE")
	r.HandleFunc("/api/wine/{WINE_ID}", updateWine).Methods("PUT")
	return r
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}

func getWine(w http.ResponseWriter, r *http.Request) {
	data.DisplayAllNotes()
	return
}

func updateWine(w http.ResponseWriter, r *http.Request) {

}

func deleteWine(w http.ResponseWriter, r *http.Request) {

}

func addWine(w http.ResponseWriter, r *http.Request) {
	var wine wine.Wine
	json.NewDecoder(r.Body).Decode(&wine)
	err := data.InsertNote(wine.Name, string(wine.Type))
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}
