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
	json.NewEncoder(w).Encode(Api.AllMenu(17))
}

func getRestaurantBanner(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	lat, err1 := strconv.ParseFloat(params["latitude"], 64)
	if err1 != nil {
		panic(err1.Error())
	}
	long, err2 := strconv.ParseFloat(params["longitude"], 64)
	if err2 != nil {
		panic(err2.Error())
	}
	json.NewEncoder(w).Encode(Api.GetRestaurantBanner(params["keyword"], lat, long))
}

func  homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Homepage")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/allmenu", allMenu).Methods("GET")
	myRouter.HandleFunc("/restaurant/{keyword}/{latitude}/{longitude}", getRestaurantBanner).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000",myRouter))
}


func main(){
	handleRequests()

}