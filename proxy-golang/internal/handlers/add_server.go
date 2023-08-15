package handlers

import (
	"encoding/json"
	"net/http"
	"proxy-golang/internal/models"

	"gorm.io/gorm"
)

func AddServer(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	var server models.Server

	err := json.NewDecoder(r.Body).Decode(&server)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return err
	}

	db.Create(&server)

	// Use the parsed JSON data
	// ...

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("JSON parsed successfully"))

	return nil
}
