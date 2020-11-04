
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

// func testPostallDetail(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Tamilarasan")

// }
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello frindes")
}

func handleRequest() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/tamil", allDetail)
	//myRouter.HandleFunc("/tamil", testPostallDetail)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	handleRequest()

}