package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/harkirath1511/mongo-api/routers"
)

func main() {
	fmt.Println("Server is getting started...")

	r := routers.Router()
	fmt.Println("Listening at port 4000!")

	r.HandleFunc("/", SayHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hi , the server's up and running!")
	fmt.Println("SERVER TEST")
}
