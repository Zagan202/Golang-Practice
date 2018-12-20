package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Display all from the people var
func get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var name1 string
	name1 = params["nm"]
	println("Hello", name1)
}

// // Display a single data
func getname(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var name1 string
	name1 = params["nm"]
	println("Hello", name1)
}

// // create a new item
// func CreatePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var person Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	person.ID = params["id"]
// 	people = append(people, person)
// 	json.NewEncoder(w).Encode(people)
// }

// // Delete an item
// func DeletePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for index, item := range people {
// 		if item.ID == params["id"] {
// 			people = append(people[:index], people[index+1:]...)
// 			break
// 		}
// 		json.NewEncoder(w).Encode(people)
// 	}
// }

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", get).Methods("GET")
	router.HandleFunc("/hello?name={nm}", getname).Methods("GET")
	// router.HandleFunc("/check", get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
