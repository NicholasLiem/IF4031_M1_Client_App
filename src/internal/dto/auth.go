package dto

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
