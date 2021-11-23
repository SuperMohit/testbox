package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


type Composer interface {
	InitContainer(domain string) (int,error)
	ParseFile(version, name, port string) (string, error)
	ApplyCompose(fileName, containerName string)(int, error)
}

type OrchestrateRouter struct {
	compose Composer
}

func NewOrchestrateRouter(compose Composer) *OrchestrateRouter {
	return &OrchestrateRouter{compose: compose}
}

func (s *OrchestrateRouter) ServiceRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/container/{domain}", s.CreateDomainInstances).Methods("POST")
	return router
}

func (s *OrchestrateRouter) CreateDomainInstances(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	domain := params["domain"]
	port, err := s.compose.InitContainer(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	fmt.Fprintf(w, strconv.Itoa(port))
}