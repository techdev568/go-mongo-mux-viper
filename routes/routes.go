package routes

import (
	"github.com/gorilla/mux"
	"github.com/parsaakbari1209/go-mongo-crud-rest-api/repository"
	"github.com/parsaakbari1209/go-mongo-crud-rest-api/service"
)

func Routes(repository repository.Repository) *mux.Router {
	// create an http server
	server := service.NewServer(repository)
	//log.Info("Router initiation")
	router := mux.NewRouter()
	//router.HandleFunc("/", server.CheckHealth).Methods("GET")

	router.HandleFunc("/users", server.CreateUser).Methods("POST")
	router.HandleFunc("/users/{email}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{email}", server.GetUser).Methods("GET")
	router.HandleFunc("/users/{email}", server.DeleteUser).Methods("DELETE")

	return router
}
