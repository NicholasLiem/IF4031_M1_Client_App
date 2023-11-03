package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"net/http"
	"time"
)

func (m *MicroserviceServer) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	_, jwtToken, err := m.authService.SignIn(loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusUnauthorized, messages.UnsuccessfulLogin)
		return
	}

	responseCookie := http.Cookie{
		Name:     "sessionData",
		Value:    jwtToken.Token,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &responseCookie)

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogin, nil)
	return
}

func (m *MicroserviceServer) Register(w http.ResponseWriter, r *http.Request) {
	var signUpDTO dto.SignupDTO
	err := json.NewDecoder(r.Body).Decode(&signUpDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if signUpDTO.Email == "" || signUpDTO.FirstName == "" || signUpDTO.LastName == "" || signUpDTO.Password == "" {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	userStruct, err := dto.SignupDTOToUserModel(signUpDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToRegister)
		return
	}

	_, err = m.authService.SignUp(userStruct)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToRegister)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulRegister, nil)
	return
}

func (m *MicroserviceServer) Logout(w http.ResponseWriter, r *http.Request) {

	expiredCookie := &http.Cookie{
		Name:     "sessionData",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, expiredCookie)

	_, err := r.Cookie("sessionData")
	if err != nil {
		response.ErrorResponse(w, http.StatusForbidden, messages.SessionExpired)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogout, nil)
}
