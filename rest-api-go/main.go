package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type details struct{
	
	Name string 
	Email string
	City  string
	State string
 
}

//var User []details

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "End point  called: homePage()")
}

func getdetailsPage(w http.ResponseWriter, r *http. Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Println("Function called : getdetailsPage()")

	//json.NewEncoder(w).Encode(getdetailsPage)

}

func handleRequest(){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",homePage).Methods("GET")
	router.HandleFunc("/mydetails", getdetailsPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",router))
}

func main(){

	    userdetails :=&details{
	 	Name : "TAmilarasan",
	 	Email : "tamilarasan@gmail.com",
	 	City : "hosur",
	 	State: "Tamilnadu",
	 }

	 _, err := json.MarshalIndent(userdetails,""," ",)
	 if err != nil{
	 	log.Println(err)
	 	os.Exit(1)
	 }

	 //fmt.Println(string(out))

	handleRequest()


}