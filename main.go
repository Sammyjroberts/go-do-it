package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeStatic(router *mux.Router, staticDirectory string) {
	staticPaths := map[string]string{
		"styles":           staticDirectory + "/styles/",
		"bower_components": staticDirectory + "/bower_components/",
		"images":           staticDirectory + "/images/",
		"scripts":          staticDirectory + "/scripts/",
	}
	fmt.Println(staticPaths)

	for pathName, pathValue := range staticPaths {
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
		fmt.Println(pathValue)
	}
}
func main() {
	router := NewRouter()

	staticDirectory := "./static"
	ServeStatic(router, staticDirectory)

	log.Fatal(http.ListenAndServe(":8080", router))
}
