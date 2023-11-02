package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := utils.VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseUserID)
		return
	}

	var newUser dto.CreateUserDTO
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if newUser.Email == "" || newUser.FirstName == "" || newUser.LastName == "" || newUser.Password == "" {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	userModel := datastruct.UserModel{
		UserID:    userID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
	}

	err = m.userService.CreateUser(userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToCreateUser)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulUserCreation, userModel)
	return
}
