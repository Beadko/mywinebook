package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Beadko/mywinebook/data"
	"github.com/Beadko/mywinebook/internal/wine"
	"github.com/gorilla/mux"
)

// AddRouterEndpoints add the actual endpoints for api

func AddRouterEndpoints(r *mux.Router) *mux.Router {
	home := homeHandler{}
	r.HandleFunc("/", home.ServeHTTP)
	r.HandleFunc("/wine", getWines).Methods("GET")
	r.HandleFunc("/wine/{id}", getWine).Methods("GET")
	r.HandleFunc("/wine", addWine).Methods("POST")
	r.HandleFunc("/wine/{id}", deleteWine).Methods("DELETE")
	// r.HandleFunc("/api/wine/{WINE_ID}", updateWine).Methods("PUT")
	return r
}

/* func sendJSONResponse(w http.ResponseWriter, data interface{}) {
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
}*/

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My wine book homepage"))
}

func getWines(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My wine list"))
	resp, err := data.GetWines()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func getWine(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My wine"))
	id := mux.Vars(r)["id"]
	resp, err := data.GetWine(id)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

/*func updateWine(w http.ResponseWriter, r *http.Request) {
	var wine wine.Wine
	json.NewDecoder(r.Body).Decode(&wine)
	data.UpdateNote(wine.Name)
}*/

func deleteWine(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DeleteWine(id)
	if err != nil {
		http.Error(w, "Failed to delete a note", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Note deleted successfully")
}

func addWine(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add wine"))
	var wine wine.Wine
	if err := json.NewDecoder(r.Body).Decode(&wine); err != nil {
		return
	}
	err := data.AddWine(wine.Name, string(wine.Type))
	if err != nil {
		http.Error(w, "Failed to add a note", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Note created successfully")
}
