package app

import "github.com/NicholasLiem/IF4031_M1_Client_App/internal/service"

type MicroserviceServer struct {
	userService    service.UserService
	authService    service.AuthService
	bookingService service.BookingService
}

func NewMicroservice(
	userService service.UserService,
	authService service.AuthService,
	bookingService service.BookingService,
) *MicroserviceServer {
	return &MicroserviceServer{
		userService:    userService,
		authService:    authService,
		bookingService: bookingService,
	}
}
