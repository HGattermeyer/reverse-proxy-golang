package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"proxy-golang/internal/config"
	"proxy-golang/internal/models"
	data "proxy-golang/internal/repository/server"
	"sync"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var serverMutex sync.Mutex

func GetProxyHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
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

	// Retrieve the RedirectServer based on strategy and build the URI
	rs, err := requestNextServer(&server)
	if err != nil {
		http.Error(w, "The strategy is not defined", http.StatusInternalServerError)
	}

	// Build the Uri
	uri, err = url.JoinPath(rs, config.PathPrefix, uri)

	if err != nil {
		http.Error(w, "The server could not be resolver", http.StatusInternalServerError)
	}

	// Updates the RedirectCount
	if server.RedirectCount >= len(server.RedirectServer) {
		server.RedirectCount = 0
	}

	println("Redirecting to:", uri)

	data.SaveOrUpdateServer(server, db)

	// Redirect with a 302 status code
	http.Redirect(w, r, uri, http.StatusFound)
	return nil
}

// Rules:
func requestNextServer(server *models.Server) (string, error) {
	var err error

	if server.Strategy == "RoundRobin" {
		return retrieveNextServer(server), nil
	}

	if server.Strategy == "Random" {
		return retrieveRandomServer(server), nil
	}

	if server.Strategy == "LeastAccessed" {
		return retrieveLeastAccessedServer(server), nil
	}

	err = fmt.Errorf("unknown strategy: %s", server.Strategy)
	return "", err
}

// Default (sequentially/round robin)
func retrieveNextServer(server *models.Server) string {
	rs := &server.RedirectServer[server.RedirectCount]

	// Update the counters
	server.RedirectServer[server.RedirectCount].AccessCount++
	server.RedirectCount++

	println("inside", server.RedirectCount)

	return rs.Server
}

// Random
func retrieveRandomServer(server *models.Server) string {
	// Generate random index
	min := 0
	max := len(server.RedirectServer) - 1

	random := rand.Intn(max-min) + min

	println(random)

	rs := &server.RedirectServer[random]

	// Update the counters
	server.RedirectCount++
	server.RedirectServer[random].AccessCount++

	return rs.Server
}

// Least accessed
func retrieveLeastAccessedServer(server *models.Server) string {
	redirectServers := server.RedirectServer

	// Find the index of the least accessed RedirectServer
	leastAccessedIndex := 0
	leastAccessCount := redirectServers[0].AccessCount
	for i, rs := range redirectServers {
		if rs.AccessCount < leastAccessCount {
			leastAccessCount = rs.AccessCount
			leastAccessedIndex = i
		}
	}

	// Increment the access count of the least accessed RedirectServer
	redirectServers[leastAccessedIndex].AccessCount++

	// Return the server URI
	return redirectServers[leastAccessedIndex].Server
}
