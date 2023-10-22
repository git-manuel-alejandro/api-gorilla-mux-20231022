package main

import (
	"api/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	prefix := "/api/v1"
	r.HandleFunc(prefix+"/", handlers.Home)
	// r.HandleFunc(prefix+"/users", handlers.Home).Methods("GET")
	r.HandleFunc(prefix+"/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc(prefix+"/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc(prefix+"/users/{id}", handlers.GetUserById).Methods("GET")
	r.HandleFunc(prefix+"/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc(prefix+"/users/{id}", handlers.DeleteUsers).Methods("DELETE")

	fmt.Println("server running ...")
	http.ListenAndServe(":3000", r)
}
