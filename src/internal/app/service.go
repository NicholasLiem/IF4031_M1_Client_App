package app

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/clients"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/service"
)

type MicroserviceServer struct {
	restClient     clients.RestClient
	userService    service.UserService
	authService    service.AuthService
	bookingService service.BookingService
}

func NewMicroservice(
	restClient clients.RestClient,
	userService service.UserService,
	authService service.AuthService,
	bookingService service.BookingService,
) *MicroserviceServer {
	return &MicroserviceServer{
		restClient:     restClient,
		userService:    userService,
		authService:    authService,
		bookingService: bookingService,
	}
}
