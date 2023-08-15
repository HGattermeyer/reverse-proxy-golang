package routes

import (
	"proxy-golang/internal/data"
	"proxy-golang/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(data.PathPrefix).Subrouter()
	subrouter.HandleFunc("/{param:.+}", handlers.ServerHandler).Methods("GET")
	subrouter.HandleFunc("/servers/upload-file", handlers.UploadFile).Methods("POST")
	subrouter.HandleFunc("/servers/add", handlers.AddServer).Methods("POST")

	return router
}
