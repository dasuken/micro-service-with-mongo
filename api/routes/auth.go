package routes

import (
	"microservices/api/resthandlers"
	"net/http"
)

func  NewAuthRoutes(authHandlers resthandlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Path: "/signup",
			Method: http.MethodPost,
			Handler: authHandlers.SignUp,
		},
		{
			Path: "/users/{id}",
			Method: http.MethodGet,
			Handler: authHandlers.GetUser,
		},
		{
			Path: "/users",
			Method: http.MethodGet,
			Handler: authHandlers.GetUsers,
		},
		{
			Path: "/users",
			Method: http.MethodPut,
			Handler: authHandlers.PutUser,
		},

		{
			Path: "/users/{id}",
			Method: http.MethodDelete,
			Handler: authHandlers.DeleteUser,
		},
	}
}