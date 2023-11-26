package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/app"
)

func BookingRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/booking",
		SubRoutes: []structs.Route{
			{
				"Create a new booking",
				"POST",
				"",
				server.CreateBooking,
				true,
			},
			{
				"Cancel a booking",
				"POST",
				"/cancel/{booking_id}",
				server.CancelBooking,
				true,
			},
			{
				"Update a booking",
				"PUT",
				"/{booking_id}",
				server.UpdateBooking,
				true,
			},
			{
				"Delete a booking",
				"DELETE",
				"/{booking_id}",
				server.DeleteBooking,
				true,
			},
			{
				"Get booking data",
				"GET",
				"/{booking_id}",
				server.GetBooking,
				true,
			},
		},
	}
}
