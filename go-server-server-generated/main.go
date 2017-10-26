package main

import (
	// sw "go-server-server-generated/go"
	sw "go-server-server-generated/go"
	application "server-application"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	// TODO: Will pass pointer to app to NewRouter...
	router := sw.NewRouter()
	
	r := application.Rectangle{"sample_rect", 3, 4}

	//app := application.Application{}
	application.App = & application.Application{}
	application.App.AddQuestion("This is a sample question...")

	log.Printf("Server started", r.Area(), *application.App)
	log.Fatal(http.ListenAndServe(":8080", router))
}
