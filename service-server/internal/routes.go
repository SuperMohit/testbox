package internal

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ServiceRouter struct {
	
}

func (s *ServiceRouter) ServiceRouter() *mux.Router {
	router := mux.NewRouter()
	// configures FileSystem, File Uploads, DB creations
	router.HandleFunc("/service/configure/{domain}", s.ConfigurationHandler).Methods("POST")
	return router
}

func (s *ServiceRouter) ConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	// TODO compelete DB creation
	// S3 upload
	//
}