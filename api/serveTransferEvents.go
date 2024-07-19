package api

import (
	"block-banter/database"
	"block-banter/repository"
	"encoding/json"
	"net/http"
)

func ServeTransferEvents(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewTransferEventRepository(database.DB)
	events, err := repo.List()
	if err != nil {
		http.Error(w, "Error fetching events from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "Error encoding events to JSON", http.StatusInternalServerError)
	}
}
