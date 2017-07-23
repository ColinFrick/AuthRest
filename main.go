package main

import (
	"log"
	"net/http"
	"github.com/NiciiA/AuthRest/routes"
)

/**
	main func;
 */
func main() {
	router := Routes.NewRouter(Routes.RouteList)

	log.Fatal(http.ListenAndServe(":8080", router))
}