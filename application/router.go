package Routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

// Routes struct
type Routes []Route

/**
	NewRouter function
	register list of routers
	add a logger to handlers
 */
func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}