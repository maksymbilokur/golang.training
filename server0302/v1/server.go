package main
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`

}
var usersarr []User
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersarr)
}
func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	usersarr = append(usersarr, user)
	json.NewEncoder(w).Encode(&user)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range usersarr {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range usersarr {
		if item.ID == params["id"] {
			usersarr = append(usersarr[:index], usersarr[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(usersarr)
}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	router := mux.NewRouter()
	apiRoute := router.PathPrefix("/api/v1").Subrouter()
	apiRoute.HandleFunc("/list", getUsers).Methods("GET")
	apiRoute.HandleFunc("/list", addUser).Methods("POST")
	apiRoute.HandleFunc("/list/{id}", getUser).Methods("GET")
	apiRoute.HandleFunc("/list/{id}", deleteUser).Methods("DELETE")
	http.ListenAndServe(":80", apiRoute)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err.Error())
	}
}