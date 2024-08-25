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
	r.HandleFunc("/wine", getWines).Methods("GET")
	r.HandleFunc("/wine/{id}", getWine).Methods("GET")
	r.HandleFunc("/wine", addWine).Methods("POST")
	r.HandleFunc("/wine/{id}", deleteWine).Methods("DELETE")
	r.HandleFunc("/wine/{id}", updateWine).Methods("PUT")
	r.HandleFunc("/wine_type", getWineTypes).Methods("GET")
	r.HandleFunc("/wine_type", addWineType).Methods("POST")
	r.HandleFunc("/wine_type/{id}", updateWineType).Methods("PUT")
	r.HandleFunc("/wine_type/{id}", deleteWineType).Methods("DELETE")
	r.HandleFunc("/country", getCountries).Methods("GET")
	r.HandleFunc("/country", addCountry).Methods("POST")
	r.HandleFunc("/country/{id}", updateCountry).Methods("PUT")
	r.HandleFunc("/country/{id}", deleteCountry).Methods("DELETE")
	r.HandleFunc("/aroma", getAromas).Methods("GET")
	r.HandleFunc("/aroma", addAroma).Methods("POST")
	r.HandleFunc("/aroma/{id}", updateAroma).Methods("PUT")
	r.HandleFunc("/aroma/{id}", deleteAroma).Methods("DELETE")
	r.HandleFunc("/flavour", getFlavours).Methods("GET")
	r.HandleFunc("/flavour", addFlavour).Methods("POST")
	r.HandleFunc("/flavour/{id}", updateFlavour).Methods("PUT")
	r.HandleFunc("/flavour/{id}", deleteFlavour).Methods("DELETE")
	r.HandleFunc("/tannin", getTannins).Methods("GET")
	r.HandleFunc("/balance", getBalances).Methods("GET")
	r.HandleFunc("/finish", getFinishes).Methods("GET")
	r.HandleFunc("/body", getBodies).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return r
}

func getWines(w http.ResponseWriter, r *http.Request) {
	winelist, err := data.GetWines()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		http.Error(w, "Failed to find the wine", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func updateWine(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	wine, err := data.GetWine(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to find the wine", http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&wine); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode updateWine input", http.StatusInternalServerError)
		return
	}
	err = data.UpdateWine(wine, id)
	if err != nil {
		log.Println(err)
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
		log.Printf("Failed to decode addWine input: %v", err)
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}
	id, err := data.AddWine(wine)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to add the wine", http.StatusInternalServerError)
		return
	}
	wine.ID = id
	wJSON, err := json.Marshal(wine)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(wJSON))
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

func getCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := data.GetCountries()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get countries", http.StatusInternalServerError)
		return
	}
	countriesJSON, err := json.Marshal(countries)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", countriesJSON)
}

func addWineType(w http.ResponseWriter, r *http.Request) {
	var wt wine.WineType
	if err := json.NewDecoder(r.Body).Decode(&wt); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode addWineType input", http.StatusInternalServerError)
		return
	}
	id, err := data.AddWineType(wt.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to add the wine type", http.StatusInternalServerError)
		return
	}
	wt.ID = id
	wtJSON, err := json.Marshal(wt)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(wtJSON))
}

func updateWineType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var wt wine.WineType
	if err := json.NewDecoder(r.Body).Decode(&wt); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode updateWineType input", http.StatusInternalServerError)
		return
	}
	err := data.UpdateWineType(id, wt.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update the wine type", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Wine type updated successfully")
}

func addCountry(w http.ResponseWriter, r *http.Request) {
	var country wine.Country
	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode addCountry input", http.StatusInternalServerError)
		return
	}
	id, err := data.AddCountry(country.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to add the country", http.StatusInternalServerError)
		return
	}
	country.ID = id
	countryJSON, err := json.Marshal(country)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(countryJSON))
}

func updateCountry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c wine.Country
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode updateCountry input", http.StatusInternalServerError)
		return
	}
	err := data.UpdateCountry(id, c.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update the country", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Country updated successfully")
}

func deleteWineType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DeleteWineType(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete the wine type", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Wine type deleted successfully")
}

func deleteCountry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DeleteCountry(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete the country", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Country deleted successfully")
}

func getAromas(w http.ResponseWriter, r *http.Request) {
	a, err := data.GetAromas()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get aromas", http.StatusInternalServerError)
		return
	}
	aJSON, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", aJSON)
}

func addAroma(w http.ResponseWriter, r *http.Request) {
	var a wine.Aroma
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode addAroma input", http.StatusInternalServerError)
		return
	}
	err := data.AddAroma(a.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to add the aroma", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Aroma added successfully")
}

func updateAroma(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a wine.Aroma
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode updateAroma input", http.StatusInternalServerError)
		return
	}
	err := data.UpdateAroma(id, a.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update the aroma", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Aroma updated successfully")
}

func deleteAroma(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DeleteAroma(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete the aroma", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Aroma deleted successfully")
}

func getFlavours(w http.ResponseWriter, r *http.Request) {
	f, err := data.GetFlavours()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get flavours", http.StatusInternalServerError)
		return
	}
	fJSON, err := json.Marshal(f)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", fJSON)
}

func addFlavour(w http.ResponseWriter, r *http.Request) {
	var f wine.Flavour
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode addFlavour input", http.StatusInternalServerError)
		return
	}
	err := data.AddFlavour(f.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to add the flavour", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Flavour added successfully")
}

func updateFlavour(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var f wine.Flavour
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		log.Println(err)
		http.Error(w, "Failed to decode updateFlavour input", http.StatusInternalServerError)
		return
	}
	err := data.UpdateFlavour(id, f.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update the flavour", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Flavour updated successfully")
}

func deleteFlavour(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DeleteFlavour(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete the flavour", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Flavour deleted successfully")
}

func getTannins(w http.ResponseWriter, r *http.Request) {
	t, err := data.GetTannins()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get tannins", http.StatusInternalServerError)
		return
	}
	tJSON, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(tJSON))
}

func getBodies(w http.ResponseWriter, r *http.Request) {
	b, err := data.GetBodies()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get bodies", http.StatusInternalServerError)
		return
	}
	bJSON, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bJSON))
}

func getFinishes(w http.ResponseWriter, r *http.Request) {
	f, err := data.GetFinishes()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get finishes", http.StatusInternalServerError)
		return
	}
	fJSON, err := json.Marshal(f)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(fJSON))
}

func getBalances(w http.ResponseWriter, r *http.Request) {
	b, err := data.GetBalances()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get balances", http.StatusInternalServerError)
		return
	}
	bJSON, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Could not not marshall to JSON.\nStopping here.", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bJSON))
}
