package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	serverSliceMutex sync.Mutex
)

// represents the server data structure
type Server struct {
	Uri             string   `json:"uri"`
	RedirectCount   int      `json:"redirectCount"`
	RedirectServers []string `json:"redirectServers"`
}

// server slice to register the servers to redirect
var serverSlice []Server

var pathPrefix = "/api/v1"

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
	targetURL := redirectServer + pathPrefix + "/" + targetUri

	fmt.Println("targetUrl:", targetURL)

	// Perform the redirection with a 302 status code
	http.Redirect(w, r, targetURL, http.StatusFound)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var payload []Server
	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	serverSlice = append(serverSlice, payload...)

	fmt.Println(serverSlice)

}

func AddServer(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var payload []Server
	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	serverSlice = append(serverSlice, payload...)

	fmt.Println(serverSlice)

}

func findRedirectServer(uri string) (string, bool) {
	// Thread safe to avoid concurrency issues
	serverSliceMutex.Lock()
	defer serverSliceMutex.Unlock()

	for i := range serverSlice {
		srv := &serverSlice[i]
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

func main() {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(pathPrefix).Subrouter()
	subrouter.HandleFunc("/{param:.+}", ServerHandler).Methods("GET")
	subrouter.HandleFunc("/servers/upload-file", UploadFile).Methods("POST")
	subrouter.HandleFunc("/servers/add", AddServer).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
