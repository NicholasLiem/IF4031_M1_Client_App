package app

import (
	"encoding/json"
	"net/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
)

func (m *MicroserviceServer) WebhookBookingHandler(w http.ResponseWriter, r *http.Request) {
	var incomingInvoicePayload dto.IncomingInvoicePayload
	err := json.NewDecoder(r.Body).Decode(&incomingInvoicePayload)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	updatedBookingData, httpError := m.bookingService.UpdateStatusBooking(incomingInvoicePayload.BookingID, incomingInvoicePayload)
	if httpError != nil {
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataUpdate, updatedBookingData)
	return
}
