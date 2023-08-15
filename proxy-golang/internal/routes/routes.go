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
	subrouter.HandleFunc("/{param:.+}", handlers.ServerHandler).Methods("GET")
	subrouter.HandleFunc("/servers/upload-file", handlers.UploadFile).Methods("POST")

	subrouter.HandleFunc("/servers/add", func(w http.ResponseWriter, r *http.Request) {
		err := handlers.AddServer(w, r, db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

	return router
}
