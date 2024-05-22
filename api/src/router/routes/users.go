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
		RequerireAuth: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Function:      controllers.UpdateUser,
		RequerireAuth: true,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Function:      controllers.DeleteUser,
		RequerireAuth: true,
	},
	{
		URI:           "/users/{userId}/follow",
		Method:        http.MethodPost,
		Function:      controllers.FollowUser,
		RequerireAuth: true,
	},	
	{
		URI:           "/users/{userId}/unfollow",
		Method:        http.MethodPost,
		Function:      controllers.UnfollowUser,
		RequerireAuth: true,
	},	
	{
		URI:           "/users/{userId}/followers",
		Method:        http.MethodGet,
		Function:      controllers.GetFollowers,
		RequerireAuth: true,
	},
}
