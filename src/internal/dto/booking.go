package dto

import "github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"

type CreateBookingDTO struct {
	CustomerID uint `json:"customer_id"`
	EventID    uint `json:"event_id"`
	SeatID     uint `json:"seat_id"`
}

type UpdateBookingDTO struct {
	CustomerID uint   `json:"customer_id"`
	EventID    uint   `json:"event_id"`
	SeatID     uint   `json:"seat_id"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

type TicketAppBookingResponseDTO struct {
	BookingID uint                     `json:"booking_id"`
	Status    datastruct.BookingStatus `json:"status"`
	Message   string                   `json:"message"`
}
