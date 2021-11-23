package internal

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type SignupRouter struct {
	repo *SignupRepo
}

func NewSignupRouter(repo *SignupRepo) *SignupRouter {
	return &SignupRouter{repo: repo}
}

func (s *SignupRouter) ServiceRouter() *mux.Router {
	router := mux.NewRouter()
	// configures FileSystem, File Uploads, DB creations
	router.HandleFunc("/service/signup", s.Signup).Methods("POST")
	return router
}

func (s *SignupRouter) Signup(w http.ResponseWriter, r *http.Request) {
	userInfo := User{}
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// save info into repository
	err  = s.repo.SaveInfo(userInfo)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pc := ProxyConfig{}

	//port, err := clients.ScheduleContainer(userInfo.SubDomain)
	//if err !=nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	err = pc.ConfigureProxyPass("44567", userInfo.SubDomain)
	if err !=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

type User struct {
	ID    		uuid.UUID  `json:"id", gorm:"primaryKey"`
	Name  		string	   `json:"name"`
	SubDomain  	string	   `json:"subDomain"`
}


