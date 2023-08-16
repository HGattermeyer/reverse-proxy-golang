package routes

import (
	"net/http"
	"proxy-golang/internal/config"
	"proxy-golang/internal/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(config.PathPrefix).Subrouter()

	// UPLOAD FILE
	subrouter.HandleFunc("/servers/upload-file", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.UploadTxtFile(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

	// SERVER CREATE
	subrouter.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.CreateServerHandler(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

	// SERVER DELETE
	subrouter.HandleFunc("/servers/{param:.+}", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.DeleteServerHandler(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("DELETE")

	// SERVER GET BY URI
	subrouter.HandleFunc("/servers/{param:.+}", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.GetServerByUriHandler(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	// SERVER GET ALL
	subrouter.HandleFunc("/servers/", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.GetAllServersHandler(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	// SERVER UPDATE STRATEGY
	subrouter.HandleFunc("/servers/{param:.+}/strategy", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.UpdateStrategyHandler(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("PUT")

	// PROXY
	subrouter.HandleFunc("/{param:.+}", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.GetProxyHandler(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	return router
}
