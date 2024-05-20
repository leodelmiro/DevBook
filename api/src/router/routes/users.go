package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Function:      controllers.CreateUser,
		RequerireAuth: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Function:      controllers.GetUsers,
		RequerireAuth: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodGet,
		Function:      controllers.GetUser,
		RequerireAuth: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Function:      controllers.UpdateUser,
		RequerireAuth: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Function:      controllers.DeleteUser,
		RequerireAuth: false,
	},
}
