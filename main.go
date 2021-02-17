package main

import (
	"encoding/json"
	Api "example.com/temporaryBackendSe/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func allMenu(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Api.AllMenu())
}


func  homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Homepage")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/allmenu", allMenu).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000",myRouter))
}


func main(){
	handleRequests()

}