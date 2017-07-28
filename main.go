package main

import (
	"log"
	"net/http"
	"github.com/NiciiA/AuthRest/application"
	"github.com/NiciiA/AuthRest/config"
)

/**
	main func;
 */
func main() {
	router := Routes.NewRouter(Routes.RouteList)

	log.Fatal(http.ListenAndServe(Config.ApplicationPort, router))
}