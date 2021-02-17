package main

import (
	"encoding/json"
	Api "example.com/temporaryBackendSe/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func allMenu(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Api.AllMenu())
}

func getRestaurant(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(Api.GetRestaurant(i))
}

func  homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Homepage")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/allmenu", allMenu).Methods("GET")
	myRouter.HandleFunc("/restaurant/{id}", getRestaurant).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000",myRouter))
}


func main(){
	handleRequests()
}