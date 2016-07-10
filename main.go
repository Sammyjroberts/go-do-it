package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	cssHandler := http.FileServer(http.Dir("./css/"))
	imagesHandler := http.FileServer(http.Dir("./images/"))
	scriptHandler := http.FileServer(http.Dir("./scripts/"))

	http.Handle("/scripts/", http.StripPrefix("/scripts/", scriptHandler))
	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/images/", http.StripPrefix("/images/", imagesHandler))
	log.Fatal(http.ListenAndServe(":8080", router))
}
