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

	json.NewEncoder(w).Encode(userdetails)
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




























/*
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


	// details := Userdetail{
	// 	User{Name: "Tamilarasan", Email: "tamil@gmail.com", Address: "Hosur"},
	// }
	

	// detail := Userdetail{
	// 	User{Name: "Vicky", Email: "vicky.com", Address: "Bengaluru"},
	// }
	
	fmt.Println("User detail")
	json.NewEncoder(w).Encode(details)
	// json.NewEncoder(w).Encode(detail)

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

	//    details = append(Userdetail, User{
    //       	Name: "Tamilarasan",
	//    	 Email: "tamil@gmail.com", 
	// 	 Address: "Hosur",
	// })
	details := Userdetail{
		User{Name: "Tamilarasan", Email: "tamil@gmail.com", Address: "Hosur"},
	}
	handleRequest()

}*/