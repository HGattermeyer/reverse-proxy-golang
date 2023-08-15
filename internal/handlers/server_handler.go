package handlers

import (
	"fmt"
	"net/http"
	"proxy-golang/internal/data"
	"sync"

	"github.com/gorilla/mux"
)

var (
	serverSliceMutex sync.Mutex
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	targetUri := params["param"]

	// Check if the uri exists on mapping
	redirectServer, exist := findRedirectServer(targetUri)
	if !exist {
		http.NotFound(w, r)
		return
	}

	// Construct the target URL for redirection
	targetURL := redirectServer + data.PathPrefix + "/" + targetUri

	fmt.Println("targetUrl:", targetURL)

	// Perform the redirection with a 302 status code
	http.Redirect(w, r, targetURL, http.StatusFound)
}

func findRedirectServer(uri string) (string, bool) {
	// Thread safe to avoid concurrency issues
	serverSliceMutex.Lock()
	defer serverSliceMutex.Unlock()

	for i := range data.ServerSlice {
		srv := &data.ServerSlice[i]
		if srv.Uri == uri && len(srv.RedirectServers) > 0 {
			srvCount := len(srv.RedirectServers)
			if srv.RedirectCount == srvCount {
				srv.RedirectCount = 0
			}

			redirectServer := srv.RedirectServers[srv.RedirectCount]

			srv.RedirectCount++
			return redirectServer, true
		}
	}
	return "", false
}
