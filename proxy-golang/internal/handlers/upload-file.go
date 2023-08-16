package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"proxy-golang/internal/models"
	data "proxy-golang/internal/repository/server"

	"gorm.io/gorm"
)

func UploadTxtFile(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	// Get the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the first line (endpoint)
	var uri string
	if scanner.Scan() {
		uri = scanner.Text()
		fmt.Println("First line:", uri)
	} else {
		http.Error(w, "Error reading first line of file", http.StatusInternalServerError)
		return scanner.Err()
	}

	// Check if the uri already exists in db
	serverExists, err := data.GetServerByUri(uri, db)

	if serverExists.ID != 0 {
		http.Error(w, "This URI already exists", http.StatusBadRequest)
		return err
	}

	// Reading the remaining lines (Redirect servers)
	var rs []string
	for scanner.Scan() {
		line := scanner.Text()
		rs = append(rs, line)
	}

	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading remaining lines of file", http.StatusInternalServerError)
		return err
	}

	// Create the RedirectServer object
	var redirectServers []models.RedirectServer
	for _, line := range rs {
		redirectServer := models.RedirectServer{
			Server: line,
		}
		redirectServers = append(redirectServers, redirectServer)
	}

	// Create the Server object
	server := models.Server{
		Uri:            uri,
		RedirectCount:  0,
		RedirectServer: redirectServers,
	}

	data.SaveOrUpdateServer(server, db)

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded and processed successfully"))
	return nil
}
