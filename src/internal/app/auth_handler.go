package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
)

func (m *MicroserviceServer) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	decodeError := json.NewDecoder(r.Body).Decode(&loginDTO)
	if decodeError != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	_, jwtToken, httpError := m.authService.SignIn(loginDTO)
	if httpError != nil {
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
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
	decodeError := json.NewDecoder(r.Body).Decode(&signUpDTO)
	if decodeError != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if signUpDTO.Email == "" || signUpDTO.FirstName == "" || signUpDTO.LastName == "" || signUpDTO.Password == "" {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	userStruct, err := dto.SignupDTOToUser(signUpDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.InvalidRequestData)
		return
	}

	_, httpError := m.authService.SignUp(userStruct)
	if httpError != nil {
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
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
