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

func postSubmitCart(w http.ResponseWriter,r *http.Request){
	var order Api.Order
	err:=json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
    }
	json.NewEncoder(w).Encode(Api.AddOrder(order))
}

func postReview(w http.ResponseWriter,r *http.Request){
	var review Api.Review
	err:=json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
    }
	Api.AddReview(review)
}

func allMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	R_id, err1 := strconv.ParseFloat(params["R_id"], 64)
	if err1 != nil {
		panic(err1.Error())
	}
	json.NewEncoder(w).Encode(Api.AllMenu(int(R_id)))
}

func getRestaurant(w http.ResponseWriter, r *http.Request) {
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

func getBanner(w http.ResponseWriter, r *http.Request) {
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

func getStatus(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    O_ID, err1 := strconv.ParseFloat(params["OrderId"], 64)
    if err1 != nil {
        panic(err1.Error())
    }
    D_ID, err2 := strconv.ParseFloat(params["DeliveryId"], 64)
    if err2 != nil {
        panic(err2.Error())
    }
	 json.NewEncoder(w).Encode(Api.OrderInfo(int(D_ID),int(O_ID)))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage hi")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allmenu/{R_id}", allMenu).Methods("GET")
	myRouter.HandleFunc("/restaurant/{keyword}/{latitude}/{longitude}", getRestaurant).Methods("GET")
	myRouter.HandleFunc("/banner/{latitude}/{longitude}", getBanner).Methods("GET")
	myRouter.HandleFunc("/submitcart", postSubmitCart).Methods("POST")
	myRouter.HandleFunc("/review", postReview).Methods("POST")
	myRouter.HandleFunc("/waitingFood/{OrderId}/{DeliveryId}", getStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", myRouter))
}

func main() {
	handleRequests()

}
