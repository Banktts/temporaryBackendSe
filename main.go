package main

import(	
	Api "example.com/temporaryBackendSe/api"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	
)

func  homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Homepage")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	log.Fatal(http.ListenAndServe(":9000",myRouter))
}


func main(){
	handleRequests()
	Api.Test()
}