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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	R_id, err1 := strconv.ParseFloat(params["R_id"], 64)
	if err1 != nil {
		panic(err1.Error())
	}
	json.NewEncoder(w).Encode(Api.AllMenu(int(R_id)))
	
}

func getRestaurant(w http.ResponseWriter, r *http.Request){
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
	json.NewEncoder(w).Encode(Api.GetRestaurant(params["keyword"], lat, long))
}

func getBanner(w http.ResponseWriter, r *http.Request){
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

	fmt.Println(lat, long)
	json.NewEncoder(w).Encode(Api.GetBanner(lat, long))
}

func  homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Homepage")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/allmenu/{R_id}", allMenu).Methods("GET")
	myRouter.HandleFunc("/restaurant/{keyword}/{latitude}/{longitude}", getRestaurant).Methods("GET")
	myRouter.HandleFunc("/banner/{latitude}/{longitude}", getBanner).Methods("GET")
	log.Fatal(http.ListenAndServe(":80",myRouter))
}


func main(){
	handleRequests()
}