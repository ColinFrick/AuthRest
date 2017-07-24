package main

import (
	"log"
	"net/http"
	"github.com/NiciiA/AuthRest/application"
)

/**
	main func;
 */
func main() {
	router := Routes.NewRouter(Routes.RouteList)

	log.Fatal(http.ListenAndServe(":8080", router))
}