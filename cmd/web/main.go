package main

import (
	"log"
	"net/http"
)

func main() {

	mux := routes()

	log.Println("Starting a web server on port 3000")

	_ = http.ListenAndServe(":3000", mux)
}
