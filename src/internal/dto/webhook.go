package dto

import uuid "github.com/satori/go.uuid"

type IncomingInvoicePayload struct {
	InvoiceID     uuid.UUID     `json:"id,omitempty"`
	BookingID     uuid.UUID     `json:"bookingID,omitempty"`
	PaymentURL    string        `json:"paymentURL,omitempty"`
	PaymentStatus PaymentStatus `json:"paymentStatus,omitempty"`
	Message       string        `json:"message,omitempty"`
}

type PaymentStatus string

const (
	SUCCESS PaymentStatus = "SUCCESS"
	PENDING PaymentStatus = "PENDING"
	FAILED  PaymentStatus = "FAILED"
)
