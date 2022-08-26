package main

import (
	"fmt"
	"log"
	"net/http"
)

// Request
// Reponse
// fun(req,res){}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the QTSolution intern!")
	fmt.Println("Endpoint Hit: homePage")
}

func Name(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My name is sheriffdeen Yusuf")
	fmt.Println("Endpoint Hit: homePage")
}

func Stack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Backend intern at QTSS")
	fmt.Println("Endpoint Hit: homePage")
}

func Tool(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Learning baskend with go lang")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", Welcome)
	http.HandleFunc("/Name", Name)
	http.HandleFunc("/Stack", Stack)
	http.HandleFunc("/Tool", Tool)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}