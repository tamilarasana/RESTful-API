
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


func creatItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content- Type", "application/json")


	var  item Item

	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)

}

func deleteItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content- Type", "application/json")

	params :=  mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUid(uid string){
	for index, item := range inventory{
		if  item.UID == uid {
			//Delete item for  slice

			inventory = append(inventory[:index], inventory[index+1:]...)
			break

		}
	}
}	

func updateItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content- Type", "application/json")

	var item Item

	json.NewDecoder(r.Body).Decode(&item)

	params :=  mux.Vars(r)

		//Delete the Item at UID 

	_deleteItemAtUid(params["uid"])

		//create   it  with new data 

	inventory = append(inventory,item)

	
	json.NewEncoder(w).Encode(inventory)

}


func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/inventory", getInventory).Methods("GET")
	myRouter.HandleFunc("/inventory", creatItem).Methods("POST")
	myRouter.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	myRouter.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")
	
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
		UID: "1",
		Name: "milk",
		Desc: "A jug  of millk",
		Price: 5.99,
	})

	inventory =append(inventory,Item{
		UID: "2",
		Name: "wine",
		Desc: "A jug  of wine",
		Price: 5.99,
	})
	
	


	handleRequests()
}