package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	initStaticServer(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
