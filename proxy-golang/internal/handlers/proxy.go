package handlers

import (
	"net/http"
	"net/url"
	"proxy-golang/internal/data"
	"sync"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var serverMutex sync.Mutex

func GetProxy(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	params := mux.Vars(r)
	uri := params["param"]

	// Retrieve the server object
	server, err := data.GetServerByUri(uri, db)

	println("Getting the server:", server.ID, server.Uri, server.RedirectCount)

	// Return error if the server is not registered
	if server.ID == 0 {
		http.Error(w, "Server not registered.", http.StatusNotFound)
		return err
	}

	// Block server object
	// Lock the server object to prevent concurrent access
	serverMutex.Lock()
	defer serverMutex.Unlock()

	// Retrieve the RedirectServer and build the URI
	rs := server.RedirectServer[server.RedirectCount].Server
	uri, err = url.JoinPath(rs, data.PathPrefix, uri)

	if err != nil {
		http.Error(w, "The server could not be resolver", http.StatusInternalServerError)
	}

	println("Redirecting to:", uri)

	// Update the counter or reset the counter
	server.RedirectCount++

	if server.RedirectCount >= len(server.RedirectServer) {
		server.RedirectCount = 0
	}

	data.SaveOrUpdateServer(server, db)

	// Redirect with a 302 status code
	http.Redirect(w, r, uri, http.StatusFound)
	return nil
}

// Rules:
// Default (sequentially/round robin)
// Random
// Least accessed
