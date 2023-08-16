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

// func ProxyHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	targetUri := params["param"]

// 	// Check if the uri exists on mapping
// 	redirectServer, exist := findRedirectServer(targetUri)
// 	if !exist {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	// Construct the target URL for redirection
// 	targetURL := redirectServer + data.PathPrefix + "/" + targetUri

// 	fmt.Println("targetUrl:", targetURL)

// 	// Perform the redirection with a 302 status code
// 	http.Redirect(w, r, targetURL, http.StatusFound)
// }

// // Look into the server object to find the next server to redirect
// func findRedirectServer(uri string) (string, bool) {
// 	// Thread safe to avoid concurrency issues
// 	serverSliceMutex.Lock()
// 	defer serverSliceMutex.Unlock()

// 	for i := range data.ServerSlice {
// 		srv := &data.ServerSlice[i]
// 		if srv.Uri == uri && len(srv.RedirectServers) > 0 {
// 			srvCount := len(srv.RedirectServers)
// 			if srv.RedirectCount == srvCount {
// 				srv.RedirectCount = 0
// 			}

// 			redirectServer := srv.RedirectServers[srv.RedirectCount]

// 			srv.RedirectCount++
// 			return redirectServer, true
// 		}
// 	}
// 	return "", false
// }
