package datastruct

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	CustomerID uint          `gorm:"customer_id" json:"customer_id,omitempty"`
	InvoiceID  uint          `gorm:"invoice_id" json:"invoice_id,omitempty"`
	PaymentURL string        `gorm:"payment_url" json:"payment_url,omitempty"`
	EventID    uint          `gorm:"column:event_id" json:"event_id,omitempty"`
	SeatID     uint          `gorm:"column:seat_id" json:"seat_id,omitempty"`
	Email      string        `gorm:"column:email" json:"email,omitempty"`
	Status     BookingStatus `gorm:"column:status" json:"status,omitempty"`
	Message    string        `gorm:"column:message" json:"message,omitempty"`
	Customer   User          `gorm:"foreignKey:CustomerID" json:"-"`
}

type BookingRequestDTO struct {
	BookingID  uint   `json:"booking_id,omitempty"`
	CustomerID uint   `json:"customer_id,omitempty"`
	EventID    uint   `json:"event_id,omitempty"`
	SeatID     uint   `json:"seat_id,omitempty"`
	Email      string `json:"email,omitempty"`
}

type BookingStatus string

const (
	BookingFailed    BookingStatus = "failed"
	BookingSuccess   BookingStatus = "success"
	BookingOnProcess BookingStatus = "on-going"
)

func IsValidStatus(status BookingStatus) bool {
	return status == BookingFailed || status == BookingSuccess || status == BookingOnProcess
}
