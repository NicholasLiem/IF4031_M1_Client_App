package app

import "github.com/NicholasLiem/IF4031_M1_Client_App/internal/service"

type MicroserviceServer struct {
	userService service.UserService
	authService service.AuthService
}

func NewMicroservice(
	userService service.UserService,
	authService service.AuthService,
) *MicroserviceServer {
	return &MicroserviceServer{
		userService: userService,
		authService: authService,
	}
}
