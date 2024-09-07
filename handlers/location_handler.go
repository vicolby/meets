package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
)

func CreateLocationHandler(w http.ResponseWriter, r *http.Request) {
	var location types.Location

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.Insert(&location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(location)
}

func DeleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	var location types.Location

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.Delete(&location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := database.GetLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}
