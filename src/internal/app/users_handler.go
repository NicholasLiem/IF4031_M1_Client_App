package app

import (
	"encoding/json"
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"github.com/gorilla/mux"
)

func (m *MicroserviceServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	_, err := utils.VerifyId(id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
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

	userModel := datastruct.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
	}

	err = m.userService.CreateUser(userModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToCreateData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataCreation, userModel)
	return
}

func (m *MicroserviceServer) GetUserData(w http.ResponseWriter, r *http.Request) {
	/**
	Checking params
	*/
	params := mux.Vars(r)
	paramsUserID := params["user_id"]
	requestedUserID, err := utils.VerifyId(paramsUserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
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
	issuerId, err := utils.VerifyId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	/**
	Process the request
	*/
	userData, err := m.userService.GetUser(requestedUserID, issuerId)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToGetData)
		return
	}
	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, userData)
	return
}

func (m *MicroserviceServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsUserID := params["user_id"]

	userID, err := utils.VerifyId(paramsUserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
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
	issuerId, err := utils.VerifyId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	var updateUser dto.UpdateUserDTO
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	updatedUser, err := m.userService.UpdateUser(userID, updateUser, issuerId)
	if err != nil || updatedUser == nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToUpdateData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataUpdate, nil)
	return
}

func (m *MicroserviceServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsUserID := params["user_id"]
	userID, err := utils.VerifyId(paramsUserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
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
	issuerId, err := utils.VerifyId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	_, err = m.userService.DeleteUser(userID, issuerId)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataDeletion, nil)
	return
}

func (m *MicroserviceServer) GetBookingsFromCustomerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customer_id"]
	requestedCustomerID, err := utils.VerifyId(customerID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	sessionUser, err := utils.ParseSessionUserFromContext(r.Context())
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToParseCookie)
		return
	}

	issuerID, err := utils.VerifyId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	bookingsData, err := m.bookingService.GetBookingsFromCustomerID(issuerID, requestedCustomerID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, bookingsData)
}
