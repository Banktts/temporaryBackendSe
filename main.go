package main

import (
	"fmt"
	"log"
	"net/http"

	Api "example.com/temporaryBackendSe/api"
	"github.com/gorilla/mux"
)

func allMenu(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(Api.AllMenu())
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allmenu", allMenu).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func main() {
	//handleRequests()
	fmt.Println("sfxdc")
	fmt.Println(Api.AllMenu(1))

	//Api.AllMenu(1)

}
