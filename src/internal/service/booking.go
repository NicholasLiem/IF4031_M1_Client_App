package service

import (
	"encoding/json"
	"errors"
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/clients"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	http2 "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"io"
	"net/http"
)

type BookingService interface {
	CreateBooking(restClient clients.RestClient, issuerID uint, booking dto.CreateBookingDTO) (*datastruct.Booking, error)
	UpdateBooking(issuerID, bookingID uint, booking dto.UpdateBookingDTO) (*datastruct.Booking, error)
	DeleteBooking(issuerID, bookingID uint) (*datastruct.Booking, error)
	GetBooking(issuerID, bookingID uint) (*datastruct.Booking, error)
	GetBookingsFromCustomerID(issuerID, customerID uint) ([]datastruct.Booking, error)
}

type bookingService struct {
	dao repository.DAO
}

func NewBookingService(dao repository.DAO) BookingService {
	return &bookingService{dao: dao}
}

func (bs *bookingService) CreateBooking(restClient clients.RestClient, issuerID uint, bookingDTO dto.CreateBookingDTO) (*datastruct.Booking, error) {
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	if userBySession.Role != datastruct.ADMIN && issuerID != bookingDTO.CustomerID {
		return nil, errors.New("user isn't authorized")
	}

	booking := datastruct.Booking{
		CustomerID: bookingDTO.CustomerID,
		EventID:    bookingDTO.EventID,
		SeatID:     bookingDTO.SeatID,
	}

	requestBody, err := json.Marshal(booking)
	if err != nil {
		return nil, err
	}

	externalAPIPath := "/book"
	response, err := restClient.Post(externalAPIPath, requestBody)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("External API request failed with status code: " + response.Status)
	}

	dataBytes, err := http2.GetJSONDataBytesFromResponse(response)
	if err != nil {
		return nil, err
	}

	var bookingResponse dto.TicketAppBookingResponseDTO
	if err := json.Unmarshal(dataBytes, &bookingResponse); err != nil {
		return nil, err
	}

	createdLocalBooking, err := bs.dao.NewBookingQuery().CreateBooking(datastruct.Booking{
		CustomerID: bookingDTO.CustomerID,
		EventID:    bookingDTO.EventID,
		SeatID:     bookingDTO.SeatID,
		Status:     bookingResponse.Status,
		Message:    bookingResponse.Message,
	})
	if err != nil {
		return nil, err
	}
	return createdLocalBooking, nil
}

func (bs *bookingService) UpdateBooking(issuerID uint, bookingID uint, bookingDTO dto.UpdateBookingDTO) (*datastruct.Booking, error) {
	var userBySession *datastruct.UserModel
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	if userBySession.Role != datastruct.ADMIN {
		return nil, errors.New("user isn't authorized")
	}

	bookingData, err := bs.dao.NewBookingQuery().GetBooking(bookingID)
	if err != nil && bookingData != nil {
		return nil, errors.New("booking not found")
	}

	updatedBookingData := datastruct.Booking{
		CustomerID: bookingDTO.CustomerID,
		EventID:    bookingDTO.EventID,
		SeatID:     bookingDTO.SeatID,
		Status:     datastruct.BookingStatus(bookingDTO.Status),
		Message:    bookingDTO.Message,
	}
	updatedBooking, err := bs.dao.NewBookingQuery().UpdateBooking(bookingID, updatedBookingData)
	return updatedBooking, err
}

func (bs *bookingService) DeleteBooking(issuerID, bookingID uint) (*datastruct.Booking, error) {
	var userBySession *datastruct.UserModel
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	if userBySession.Role != datastruct.ADMIN {
		return nil, errors.New("user isn't authorized")
	}

	bookingData, err := bs.dao.NewUserQuery().GetUser(bookingID)
	if err != nil && bookingData != nil {
		return nil, errors.New("booking not found")
	}
	deletedBooking, err := bs.dao.NewBookingQuery().DeleteBooking(bookingID)
	return deletedBooking, err
}

func (bs *bookingService) GetBooking(issuerID, bookingID uint) (*datastruct.Booking, error) {
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	booking, err := bs.dao.NewBookingQuery().GetBooking(bookingID)
	if err != nil {
		return nil, err
	}

	if userBySession.Role != datastruct.ADMIN {
		if issuerID != booking.CustomerID {
			return nil, errors.New("user isn't authorized")
		}
	}

	return booking, nil
}

func (bs *bookingService) GetBookingsFromCustomerID(issuerID, customerID uint) ([]datastruct.Booking, error) {
	userBySession, err := bs.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	if userBySession.Role != datastruct.ADMIN {
		if issuerID != customerID {
			return nil, errors.New("user isn't authorized")
		}
	}

	bookings, err := bs.dao.NewBookingQuery().GetBookingsFromCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}
