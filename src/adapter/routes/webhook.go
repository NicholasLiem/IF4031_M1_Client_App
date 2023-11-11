package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/app"
)

// TODO: Implement this please
func WebhookRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/webhook",
		SubRoutes: []structs.Route{
			{
				"Update booking if failed or success",
				"PUT",
				"/",
				server.CreateBooking,
				true,
			},
		},
	}
}
