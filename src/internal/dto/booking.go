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
	InvoiceID  string                   `json:"invoice_id,omitempty"`
	BookingID  uint                     `json:"booking_id,omitempty"`
	EventID    uint                     `json:"event_id,omitempty"`
	SeatID     uint                     `json:"seat_id,omitempty"`
	CustomerID uint                     `json:"customer_id,omitempty"`
	PaymentURL string                   `json:"payment_uRL,omitempty"`
	Status     datastruct.BookingStatus `json:"status,omitempty"`
	Message    string                   `json:"message,omitempty"`
}
