package app

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsUserID := params["user_id"]
	userID, err := utils.VerifyUserId(paramsUserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseUserID)
		return
	}

	/**
	Parsing Session Data from Context
	*/
	sessionUser, err := utils.ParseSessionUserFromContext(r.Context())
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToParseCookie)
		return
	}

	/**
	Took the issuer identifier
	*/
	issuerId, err := utils.VerifyUserId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseUserID)
		return
	}

	userData, err := m.userService.DeleteUser(userID, issuerId)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToDeleteUser)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulUserDeletion, userData)
	return
}
