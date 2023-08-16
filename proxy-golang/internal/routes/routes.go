package routes

import (
	"net/http"
	"proxy-golang/internal/data"
	"proxy-golang/internal/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(data.PathPrefix).Subrouter()
	subrouter.HandleFunc("/servers/upload-file", handlers.UploadFile).Methods("POST")

	// SERVER CREATE
	subrouter.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.CreateServer(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

	// SERVER DELETE
	subrouter.HandleFunc("/servers/{param:.+}", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.DeleteServer(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("DELETE")

	// SERVER GET BY URI
	subrouter.HandleFunc("/servers/{param:.+}", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.GetServerByUri(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	// SERVER GET ALL
	subrouter.HandleFunc("/servers/", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.GetAllServers(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")
	// GET
	// PROXY
	subrouter.HandleFunc("/{param:.+}", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.GetProxy(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	return router
}
