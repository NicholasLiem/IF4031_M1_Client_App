package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/app"
)

func AuthRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/auth",
		SubRoutes: []structs.Route{
			{
				"Login",
				"POST",
				"/login",
				server.Login,
				false,
			},
			{
				"Register",
				"POST",
				"/register",
				server.Register,
				false,
			},
			{
				"Logout",
				"POST",
				"/logout",
				server.Logout,
				true,
			},
		},
	}
}
