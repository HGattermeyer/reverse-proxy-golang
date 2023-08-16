package handlers

import (
	"encoding/json"
	"net/http"
	"proxy-golang/internal/data"
	"proxy-golang/internal/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateServer(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	var server models.Server

	// Decode Json
	err := json.NewDecoder(r.Body).Decode(&server)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return err
	}

	// Check if the uri already exists in db
	serverExists, err := data.GetServerByUri(server.Uri, db)

	if serverExists.ID != 0 {
		http.Error(w, "This URI already exists", http.StatusBadRequest)
		return err
	}

	// Force the counter to be 0
	server.RedirectCount = 0

	// Create the object on DB
	server = data.SaveOrUpdateServer(server, db)

	// Respond with a success message and the server obect
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(server)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return err
	}

	return nil
}

func DeleteServer(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	params := mux.Vars(r)
	uri := params["param"]

	// Check if the uri exists in db
	server, err := data.GetServerByUri(uri, db)

	if server.ID == 0 {
		http.Error(w, "This URI does not exist", http.StatusBadRequest)
		return err
	}

	// Delete from DB
	data.DeleteServerByUri(server.Uri, db)

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))

	return nil
}

func GetServerByUri(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	// Similar to GetProxy but return json and all redirect servers without redirect
	params := mux.Vars(r)
	uri := params["param"]

	// Retrieve the server object
	server, err := data.GetServerByUri(uri, db)

	// Respond with a success message and the server obect
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(server)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return err
	}

	return nil
}

func GetAllServers(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	// Retrieve the servers object
	server, err := data.GetAllServers(db)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	// Respond with a success message and the server obect
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(server)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return err
	}

	return nil
}

//Add
//Update
//Delete
