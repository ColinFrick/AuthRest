package Routes

import (
	"github.com/NiciiA/AuthRest/routes/handler/auth"
)

// list of routes
var RouteList = Routes{
	Route{
		"Auth",
		"POST",
		"api/v1.0/auth/login",
		AuthHandler.HandleLogin,
	},
	Route{
		"Signup",
		"POST",
		"api/v1.0/auth/signup",
		AuthHandler.HandleSignup,
	},
	Route{
		"Verify",
		"POST",
		"api/v1.0/auth/verify",
		AuthHandler.HandleVerify,
	},
	/*
	Route{
		"AuthVerify",
		"GET",
		"/auth/{code}",
		AuthVerify,
	},
	Route{
		"CustomerCreate",
		"POST",
		"/customers",
		CustomerCreate,
	},
	Route{
		"CustomerUpdate",
		"PUT",
		"/customers/{customerId}",
		CustomerUpdate,
	},
	*/
}