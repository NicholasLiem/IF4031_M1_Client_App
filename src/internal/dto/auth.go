package dto

import "github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"

type LoginDTO struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignupDTO struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

func SignupDTOToUser(signupDTO SignupDTO) (datastruct.User, error) {
	userModel := datastruct.User{
		FirstName: signupDTO.FirstName,
		LastName:  signupDTO.LastName,
		Email:     signupDTO.Email,
		Password:  signupDTO.Password,
	}
	return userModel, nil
}
