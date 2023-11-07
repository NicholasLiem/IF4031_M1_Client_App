package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/clients"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	http2 "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
)

type BookingService interface {
	CreateBooking(restClient clients.RestClient, issuerID uint, booking dto.CreateBookingDTO) (*dto.IncomingBookingResponseDTO, *utils.HttpError)
	UpdateBooking(issuerID, bookingID uint, booking dto.UpdateBookingDTO) (*datastruct.Booking, *utils.HttpError)
	DeleteBooking(issuerID, bookingID uint) (*datastruct.Booking, *utils.HttpError)
	GetBooking(issuerID, bookingID uint) (*datastruct.Booking, *utils.HttpError)
	GetBookingsFromCustomerID(issuerID, customerID uint) ([]datastruct.Booking, *utils.HttpError)
}

type bookingService struct {
	dao repository.DAO
}

func NewBookingService(dao repository.DAO) BookingService {
	return &bookingService{dao: dao}
}

func (bs *bookingService) CreateBooking(restClient clients.RestClient, issuerID uint, bookingDTO dto.CreateBookingDTO) (*dto.IncomingBookingResponseDTO, *utils.HttpError) {
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if userBySession.Role != datastruct.ADMIN && issuerID != bookingDTO.CustomerID {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if bookingDTO.CustomerID == 0 {
		bookingDTO.CustomerID = issuerID
	}

	customer, err := bs.dao.NewUserQuery().GetUser(bookingDTO.CustomerID)
	if err != nil || customer == nil {
		return nil, &utils.HttpError{
			Message:    "customer not found",
			StatusCode: http.StatusNotFound,
		}
	}

	//implement rollback somehow
	newBooking, err := bs.dao.NewBookingQuery().CreateBooking(datastruct.Booking{
		CustomerID: bookingDTO.CustomerID,
		EventID:    bookingDTO.EventID,
		SeatID:     bookingDTO.SeatID,
		Email:      customer.Email,
	})
	if err != nil {
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

		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, &utils.HttpError{
			Message:    "External API request failed with status code: " + response.Status,
			StatusCode: http.StatusInternalServerError,
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

func (bs *bookingService) UpdateBooking(issuerID uint, bookingID uint, bookingDTO dto.UpdateBookingDTO) (*datastruct.Booking, *utils.HttpError) {
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
	return updatedBooking, &utils.HttpError{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func (bs *bookingService) DeleteBooking(issuerID, bookingID uint) (*datastruct.Booking, *utils.HttpError) {
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

	bookingData, err := bs.dao.NewUserQuery().GetUser(bookingID)
	if err != nil && bookingData != nil {
		return nil, &utils.HttpError{
			Message:    "booking not found",
			StatusCode: http.StatusNotFound,
		}
	}
	deletedBooking, err := bs.dao.NewBookingQuery().DeleteBooking(bookingID)
	return deletedBooking, &utils.HttpError{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func (bs *bookingService) GetBooking(issuerID, bookingID uint) (*datastruct.Booking, *utils.HttpError) {
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

	return booking, nil
}

func (bs *bookingService) GetBookingsFromCustomerID(issuerID, customerID uint) ([]datastruct.Booking, *utils.HttpError) {
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

	return bookings, nil
}
