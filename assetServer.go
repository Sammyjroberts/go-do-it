package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeStatic(router *mux.Router, staticDirectory string) {
	staticPaths := map[string]string{
		"styles":           staticDirectory + "/styles/",
		"bower_components": staticDirectory + "/bower_components/",
		"images":           staticDirectory + "/images/",
		"scripts":          staticDirectory + "/scripts/",
		"build":            staticDirectory + "/build/",
	}
	fmt.Println(staticPaths)

	for pathName, pathValue := range staticPaths {
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
		fmt.Println(pathValue)
	}
}
func initStaticServer(router *mux.Router) {
	staticDirectory := "./static"
	ServeStatic(router, staticDirectory)
}
