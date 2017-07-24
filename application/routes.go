package Routes

import (
	"github.com/NiciiA/AuthRest/application/handler/auth"
	"github.com/NiciiA/AuthRest/application/handler/middleware"
	"github.com/NiciiA/AuthRest/application/handler/users"
	"net/http"
)

// list of routes
var RouteList = Routes{
	Route{
		"Login",
		"POST",
		"/api/v1.0/auth/login",
		AuthHandler.HandleLogin,
	},
	Route{
		"Signup",
		"POST",
		"/api/v1.0/auth/signup",
		AuthHandler.HandleSignup,
	},
	Route{
		"Verify",
		"GET",
		"/api/v1.0/auth/verify",
		AuthHandler.HandleVerify,
	},
	Route{
		"UserList",
		"GET",
		"/api/v1.0/users",
		Middleware.AdminMiddleware(http.HandlerFunc(UserHandler.List)),
	},
	Route{
		"UserSingle",
		"GET",
		"/api/v1.0/users/{id:[a-f0-9]{24}}",
		Middleware.LoggedinMiddleware(http.HandlerFunc(UserHandler.Single)),
	},
	Route{
		"UserDelete",
		"DELETE",
		"/api/v1.0/users/{id:[a-f0-9]{24}}",
		Middleware.LoggedinMiddleware(http.HandlerFunc(UserHandler.Delete)),
	},
	Route{
		"UserUpdate",
		"PUT",
		"/api/v1.0/users/{id:[a-f0-9]{24}}",
		Middleware.LoggedinMiddleware(http.HandlerFunc(UserHandler.Update)),
	},
	/*
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