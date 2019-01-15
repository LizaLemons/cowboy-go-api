package main

// mux routes our requests
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Cowboy struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var cowboys []Cowboy

// show me all cowboys
func GetCowboys(w http.ResponseWriter, r *http.Request) {
	log.Println(cowboys)
	json.NewEncoder(w).Encode(cowboys)
}

// show me one cowboy
func GetCowboy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range cowboys {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Cowboy{})
}

// make me a cowboy, clown
func CreateCowboy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var cowboy Cowboy
	_ = json.NewDecoder(r.Body).Decode(&cowboy)
	cowboy.ID = params["id"]
	cowboys = append(cowboys, cowboy)
	json.NewEncoder(w).Encode(cowboys)
}

// delete me a cowboy
func DeleteCowboy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range cowboys {
		if item.ID == params["id"] {
			cowboys = append(cowboys[:index], cowboys[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cowboys)
}

func main() {
	router := mux.NewRouter()

	cowboys = append(cowboys, Cowboy{ID: "1", Firstname: "Billy", Lastname: "The Kid", Address: &Address{City: "New York City", State: "NY"}})
	cowboys = append(cowboys, Cowboy{ID: "2", Firstname: "Will", Lastname: "Rogers", Address: &Address{City: "Oologah", State: "Oklahoma"}})
	cowboys = append(cowboys, Cowboy{ID: "3", Firstname: "Jesse", Lastname: "James", Address: &Address{City: "Kearney", State: "Missouri"}})

	router.HandleFunc("/cowboys", GetCowboys).Methods("GET")
	router.HandleFunc("/cowboys/{id}", GetCowboy).Methods("GET")
	router.HandleFunc("/cowboys/{id}", CreateCowboy).Methods("POST")
	router.HandleFunc("/cowboys/{id}", DeleteCowboy).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
