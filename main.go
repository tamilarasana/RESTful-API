
package main

import (
	
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Item struct{

	UID string //'json:"UID"'
	Name string //'json:"Name"'
	Desc string //'json:"Desc"'
	Price float64 //'json:"Price"'
}

var inventory []Item



func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "End point  called: homePage()")
}

func getInventory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content- Type", "application/json")
	fmt.Println("Function Called: getInventory()")

	json.NewEncoder(w).Encode(inventory)
}


func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/inventory", getInventory).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main(){

	inventory =append(inventory,Item{
		UID: "0",
		Name: "Chesse",
		Desc: "A find block of cheese",
		Price: 4.99,
	})
	
	inventory =append(inventory,Item{
		UID: "2",
		Name: "milk",
		Desc: "A jug  of millk",
		Price: 5.99,
	})
	


	handleRequests()
}