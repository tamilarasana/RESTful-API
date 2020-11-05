
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type User struct {
	Name    string
	Email   string
	Address string
}

type Userdetail []User

func allDetail(w http.ResponseWriter, r *http.Request) {
	details := Userdetail{
		User{Name: "Tamilarasan", Email: "tamil@gmail.com", Address: "Hosur"},
	}
	fmt.Println("User detail")
	json.NewEncoder(w).Encode(details)

}

func testPostallDetail(w http.ResponseWriter, r *http.Request) {
 	fmt.Fprint(w, "Hello Tamilarasan")

 }
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello frindes")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/tamil", allDetail).Methods("GET")
	myRouter.HandleFunc("/tamil", testPostallDetail).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main() {
	handleRequest()

}