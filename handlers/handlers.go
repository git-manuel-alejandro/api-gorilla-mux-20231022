package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Txt string `json:"txt"`
}
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Message{Txt: "hola manuel"})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(Message{Txt: "something was wrong"})
		return

	}
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	json.NewEncoder(w).Encode(Message{Txt: "UpdateUser : " + id})

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Author", "manuel alejandro roa ojeda")
	w.WriteHeader(http.StatusBadRequest)
	// w.Header().Add("Content-Length", "29999")
	params := mux.Vars(r)
	id := params["id"]
	json.NewEncoder(w).Encode(Message{Txt: "GetUserById : " + id})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Txt: "GetUsers"})
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	json.NewEncoder(w).Encode(Message{Txt: "DeleteUsers :  " + id})
}
