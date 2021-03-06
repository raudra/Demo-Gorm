package routes

import (
	"net/http"
	"demo_gorm/app/controllers"
)

type Route struct{
	Name 			string
	Method 			string
	Pattern 		string
	HandlerFunc     http.HandlerFunc
}

type Routes []Route


var routes =Routes{
	Route{
		"Controller: EmployeeController      Action: Fetch",
		"GET",
		"/fetch",
		controllers.EmployeeController.Fetch,
	},
	Route{
		"Controller: FeatureController      Action: Index",
		"GET",
		"/features",
		controllers.FeatureController.Index,
	},
	Route{
		"Controller: FeatureController      Action: Create",
		"POST",
		"/features",
		controllers.FeatureController.Create,	
	},
}