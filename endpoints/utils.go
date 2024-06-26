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

func AddRouterEndpoints(r *mux.Router) *mux.Router {
	home := homeHandler{}
	r.HandleFunc("/", home.ServeHTTP)
	r.HandleFunc("/wine", getWines).Methods("GET")
	r.HandleFunc("/wine/{id}", getWine).Methods("GET")
	r.HandleFunc("/wine", addWine).Methods("POST")
	r.HandleFunc("/wine/{id}", deleteWine).Methods("DELETE")
	r.HandleFunc("/wine/{id}", updateWine).Methods("PUT")
	r.HandleFunc("/wine_type", getWineTypes).Methods("GET")
	return r
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My wine book homepage"))
}

func getWines(w http.ResponseWriter, r *http.Request) {
	winelist, err := data.GetWines()
	if err != nil {
		http.Error(w, "Failed to get wines", http.StatusInternalServerError)
		return
	}
	winelistJson, err := json.Marshal(winelist)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", winelistJson)
}

func getWine(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	resp, err := data.GetWine(id)
	if err != nil {
		http.Error(w, "Failed to find the wine", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func updateWine(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var wine wine.Wine
	if err := json.NewDecoder(r.Body).Decode(&wine); err != nil {
		return
	}
	_, err := data.UpdateWine(id, wine.Name, wine.Type.ID)
	if err != nil {
		http.Error(w, "Failed to update the wine", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Wine updated successfully")
}

func deleteWine(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DeleteWine(id)
	if err != nil {
		http.Error(w, "Failed to delete the wine", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Wine deleted successfully")
}

func addWine(w http.ResponseWriter, r *http.Request) {
	var wine wine.Wine
	if err := json.NewDecoder(r.Body).Decode(&wine); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode addWine input", http.StatusInternalServerError)
		return
	}
	err := data.AddWine(wine.Name, wine.TypeID)
	if err != nil {
		http.Error(w, "Failed to add the wine", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Wine created successfully")
}

func getWineTypes(w http.ResponseWriter, r *http.Request) {
	winetypes, err := data.GetWineTypes()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get wine types", http.StatusInternalServerError)
		return
	}
	wineTypesJSON, err := json.Marshal(winetypes)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", wineTypesJSON)
}
