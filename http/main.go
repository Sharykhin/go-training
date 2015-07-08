package main

import (
	"go-training/http/controllers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", controllers.Hello)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
