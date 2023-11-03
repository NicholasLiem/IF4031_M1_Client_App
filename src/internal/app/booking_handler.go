package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking dto.CreateBookingDTO
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if newBooking.EventID == 0 || newBooking.SeatID == 0 {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	bookingData, err := m.bookingService.CreateBooking(m.restClient, newBooking)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataCreation, bookingData)
	return
}

func (m *MicroserviceServer) GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingID := params["booking_id"]
	requestedBookingID, err := utils.VerifyId(bookingID)
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

	bookingData, err := m.bookingService.GetBooking(issuerId, requestedBookingID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, bookingData)
	return
}

func (m *MicroserviceServer) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingID := params["booking_id"]
	requestedBookingID, err := utils.VerifyId(bookingID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	var updatedBooking dto.UpdateBookingDTO
	err = json.NewDecoder(r.Body).Decode(&updatedBooking)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
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

	_, err = m.bookingService.UpdateBooking(issuerId, requestedBookingID, updatedBooking)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataUpdate, nil)
	return
}

func (m *MicroserviceServer) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingID := params["booking_id"]
	requestedBookingID, err := utils.VerifyId(bookingID)
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

	_, err = m.bookingService.DeleteBooking(issuerId, requestedBookingID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataDeletion, nil)
	return
}
