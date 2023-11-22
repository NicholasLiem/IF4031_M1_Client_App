package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/clients"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	http2 "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	uuid "github.com/satori/go.uuid"
)

type BookingService interface {
	CreateBooking(restClient clients.RestClient, issuerID uint, booking dto.CreateBookingDTO) (*dto.IncomingBookingResponseDTO, *utils.HttpError)
	UpdateBooking(issuerID uint, bookingID uuid.UUID, booking dto.UpdateBookingDTO) (*datastruct.BookingResponse, *utils.HttpError)
	DeleteBooking(issuerID uint, bookingID uuid.UUID) (*datastruct.Booking, *utils.HttpError)
	GetBooking(issuerID uint, bookingID uuid.UUID) (*datastruct.BookingResponse, *utils.HttpError)
	GetBookingsFromCustomerID(issuerID uint, customerID uint) ([]datastruct.BookingResponse, *utils.HttpError)
}

type bookingService struct {
	dao repository.DAO
}

func NewBookingService(dao repository.DAO) BookingService {
	return &bookingService{dao: dao}
}

func (bs *bookingService) CreateBooking(restClient clients.RestClient, issuerID uint, bookingDTO dto.CreateBookingDTO) (*dto.IncomingBookingResponseDTO, *utils.HttpError) {
	// Transaction query with rollback
	tx := bs.dao.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if bookingDTO.CustomerID == 0 {
		bookingDTO.CustomerID = issuerID
	}

	if userBySession.Role != datastruct.ADMIN && issuerID != bookingDTO.CustomerID {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	customer, err := bs.dao.NewUserQuery().GetUser(bookingDTO.CustomerID)
	if err != nil || customer == nil {
		return nil, &utils.HttpError{
			Message:    "customer not found",
			StatusCode: http.StatusNotFound,
		}
	}

	// Use transaction object (tx)
	newBooking, err := bs.dao.NewBookingQuery().CreateBooking(tx, datastruct.Booking{
		CustomerID: bookingDTO.CustomerID,
		EventID:    bookingDTO.EventID,
		SeatID:     bookingDTO.SeatID,
		Email:      customer.Email,
	})
	if err != nil {
		tx.Rollback()
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	booking := datastruct.BookingRequestDTO{
		BookingID:  newBooking.ID,
		CustomerID: bookingDTO.CustomerID,
		EventID:    bookingDTO.EventID,
		SeatID:     bookingDTO.SeatID,
		Email:      newBooking.Email,
	}

	requestBody, err := json.Marshal(booking)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	externalAPIPath := "/book"
	response, err := restClient.Post(externalAPIPath, requestBody)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			tx.Rollback()
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		tx.Rollback()

		// Read the response body to get the message
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, &utils.HttpError{
				Message:    "Failed to read response body",
				StatusCode: http.StatusInternalServerError,
			}
		}

		responseBody := string(bodyBytes)

		var response http2.Response
		if err := json.Unmarshal([]byte(responseBody), &response); err != nil {
			return nil, &utils.HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return nil, &utils.HttpError{
			Message:    "External API request failed with message: " + response.Message,
			StatusCode: response.StatusCode,
		}
	}

	dataBytes, err := http2.GetJSONDataBytesFromResponse(response)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	var bookingResponse dto.IncomingBookingResponseDTO
	if err := json.Unmarshal(dataBytes, &bookingResponse); err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	// Commit the transaction
	tx.Commit()

	updatedBookingData := datastruct.Booking{
		CustomerID: newBooking.CustomerID,
		EventID:    newBooking.EventID,
		SeatID:     newBooking.SeatID,
		Email:      customer.Email,
		Status:     bookingResponse.Status,
		Message:    bookingResponse.Message,
	}
	_, err = bs.dao.NewBookingQuery().UpdateBooking(newBooking.ID, updatedBookingData)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &bookingResponse, nil
}

func (bs *bookingService) UpdateBooking(issuerID uint, bookingID uuid.UUID, bookingDTO dto.UpdateBookingDTO) (*datastruct.BookingResponse, *utils.HttpError) {
	var userBySession *datastruct.User
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if userBySession.Role != datastruct.ADMIN {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	bookingData, err := bs.dao.NewBookingQuery().GetBooking(bookingID)
	if err != nil && bookingData != nil {
		return nil, &utils.HttpError{
			Message:    "booking not found",
			StatusCode: http.StatusNotFound,
		}
	}

	updatedBookingData := datastruct.Booking{
		CustomerID: bookingDTO.CustomerID,
		EventID:    bookingDTO.EventID,
		SeatID:     bookingDTO.SeatID,
		Status:     datastruct.BookingStatus(bookingDTO.Status),
		Message:    bookingDTO.Message,
	}
	updatedBooking, err := bs.dao.NewBookingQuery().UpdateBooking(bookingID, updatedBookingData)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	responseData := datastruct.BookingResponse{
		ID:         updatedBooking.ID,
		CustomerID: updatedBooking.CustomerID,
		InvoiceID:  updatedBooking.InvoiceID,
		PaymentURL: updatedBooking.PaymentURL,
		EventID:    updatedBooking.EventID,
		SeatID:     updatedBooking.SeatID,
		Email:      updatedBooking.Email,
		Status:     updatedBooking.Status,
		Message:    updatedBooking.Message,
	}

	return &responseData, nil
}

func (bs *bookingService) DeleteBooking(issuerID uint, bookingID uuid.UUID) (*datastruct.Booking, *utils.HttpError) {
	var userBySession *datastruct.User
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if userBySession.Role != datastruct.ADMIN {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	bookingData, err := bs.dao.NewBookingQuery().GetBooking(bookingID)
	if err != nil && bookingData != nil {
		return nil, &utils.HttpError{
			Message:    "booking not found",
			StatusCode: http.StatusNotFound,
		}
	}
	deletedBooking, err := bs.dao.NewBookingQuery().DeleteBooking(bookingID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return deletedBooking, nil
}

func (bs *bookingService) GetBooking(issuerID uint, bookingID uuid.UUID) (*datastruct.BookingResponse, *utils.HttpError) {
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	booking, err := bs.dao.NewBookingQuery().GetBooking(bookingID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if userBySession.Role != datastruct.ADMIN {
		if issuerID != booking.CustomerID {
			return nil, &utils.HttpError{
				Message:    "unauthorized",
				StatusCode: http.StatusUnauthorized,
			}
		}
	}

	if booking.CustomerID == userBySession.ID || userBySession.Role == datastruct.ADMIN {
		responseData := datastruct.BookingResponse{
			ID:         booking.ID,
			CustomerID: booking.CustomerID,
			InvoiceID:  booking.InvoiceID,
			PaymentURL: booking.PaymentURL,
			EventID:    booking.EventID,
			SeatID:     booking.SeatID,
			Email:      booking.Email,
			Status:     booking.Status,
			Message:    booking.Message,
		}

		return &responseData, nil
	} else {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}
}

func (bs *bookingService) GetBookingsFromCustomerID(issuerID uint, customerID uint) ([]datastruct.BookingResponse, *utils.HttpError) {
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if userBySession.Role != datastruct.ADMIN {
		if issuerID != customerID {
			return nil, &utils.HttpError{
				Message:    "unauthorized",
				StatusCode: http.StatusUnauthorized,
			}
		}
	}

	bookings, err := bs.dao.NewBookingQuery().GetBookingsFromCustomerID(customerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	var responseData []datastruct.BookingResponse

	// Map each booking to BookingResponse
	for _, booking := range bookings {
		response := datastruct.BookingResponse{
			ID:         booking.ID,
			CustomerID: booking.CustomerID,
			InvoiceID:  booking.InvoiceID,
			PaymentURL: booking.PaymentURL,
			EventID:    booking.EventID,
			SeatID:     booking.SeatID,
			Email:      booking.Email,
			Status:     booking.Status,
			Message:    booking.Message,
		}
		responseData = append(responseData, response)
	}

	return responseData, nil
}
