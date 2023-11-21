package app

import (
	"encoding/json"
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"github.com/gorilla/mux"
)

func (m *MicroserviceServer) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking dto.CreateBookingDTO
	decodeError := json.NewDecoder(r.Body).Decode(&newBooking)
	if decodeError != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if newBooking.CustomerID == 0 || newBooking.EventID == 0 || newBooking.SeatID == 0 {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	/**
	Parsing Session Data from Context
	*/
	sessionUser, httpError := utils.ParseSessionUserFromContext(r.Context())
	if httpError != nil {
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
		return
	}

	/**
	Took the issuer identifier
	*/
	issuerId, httpError := utils.VerifyId(sessionUser.UserID)
	if httpError != nil {
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
		return
	}

	bookingData, httpError := m.bookingService.CreateBooking(m.restClient, issuerId, newBooking)
	if httpError != nil {
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataCreation, bookingData)
	return
}

func (m *MicroserviceServer) GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingID := params["booking_id"]
	requestedBookingID, err := utils.VerifyUUID(bookingID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	/**
	Parsing Session Data from Context
	*/
	sessionUser, err := utils.ParseSessionUserFromContext(r.Context())
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	/**
	Took the issuer identifier
	*/
	issuerId, err := utils.VerifyId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	bookingData, err := m.bookingService.GetBooking(issuerId, requestedBookingID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, bookingData)
	return
}

func (m *MicroserviceServer) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingID := params["booking_id"]
	requestedBookingID, err := utils.VerifyUUID(bookingID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	var updatedBooking dto.UpdateBookingDTO
	decodeError := json.NewDecoder(r.Body).Decode(&updatedBooking)
	if decodeError != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	/**
	Parsing Session Data from Context
	*/
	sessionUser, err := utils.ParseSessionUserFromContext(r.Context())
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	/**
	Took the issuer identifier
	*/
	issuerId, err := utils.VerifyId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	updatedBookingData, err := m.bookingService.UpdateBooking(issuerId, requestedBookingID, updatedBooking)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataUpdate, updatedBookingData)
	return
}

func (m *MicroserviceServer) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingID := params["booking_id"]
	requestedBookingID, err := utils.VerifyUUID(bookingID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	/**
	Parsing Session Data from Context
	*/
	sessionUser, err := utils.ParseSessionUserFromContext(r.Context())
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	/**
	Took the issuer identifier
	*/
	issuerId, err := utils.VerifyId(sessionUser.UserID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	_, err = m.bookingService.DeleteBooking(issuerId, requestedBookingID)
	if err != nil {
		response.ErrorResponse(w, err.StatusCode, err.Message)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataDeletion, nil)
	return
}
