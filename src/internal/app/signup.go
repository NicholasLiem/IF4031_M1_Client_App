package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"net/http"
)

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

	userData, err := m.authService.SignUp(userStruct)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToRegister)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulRegister, userData)
	return
}
