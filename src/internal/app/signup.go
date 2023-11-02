package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) Register(w http.ResponseWriter, r *http.Request) {
	var userModel datastruct.UserModel
	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if userModel.Email == "" || userModel.FirstName == "" || userModel.LastName == "" || userModel.Password == "" {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	userData, err := m.authService.SignUp(userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToRegister)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulRegister, userData)
	return
}
