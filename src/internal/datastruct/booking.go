package datastruct

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CustomModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Booking struct {
	CustomModel
	CustomerID uint          `gorm:"customer_id" json:"customer_id,omitempty"`
	InvoiceID  uuid.UUID     `gorm:"invoice_id" json:"invoice_id,omitempty"`
	PaymentURL string        `gorm:"payment_url" json:"payment_url,omitempty"`
	EventID    uint          `gorm:"column:event_id" json:"event_id,omitempty"`
	SeatID     uint          `gorm:"column:seat_id" json:"seat_id,omitempty"`
	Email      string        `gorm:"column:email" json:"email,omitempty"`
	Status     BookingStatus `gorm:"column:status" json:"status,omitempty"`
	Message    string        `gorm:"column:message" json:"message,omitempty"`
	Customer   User          `gorm:"foreignKey:CustomerID" json:"-"`
}

type BookingRequestDTO struct {
	BookingID  uuid.UUID `json:"booking_id,omitempty"`
	CustomerID uint      `json:"customer_id,omitempty"`
	EventID    uint      `json:"event_id,omitempty"`
	SeatID     uint      `json:"seat_id,omitempty"`
	Email      string    `json:"email,omitempty"`
}

type CancelBookingRequest struct{
	BookingID  uuid.UUID `json:"booking_id,omitempty"`
	SeatID     uint      `json:"seat_id,omitempty"`
}

type CancelBookingResponse struct{
	BookingID	uuid.UUID	`json:"booking_id,omitempty"`
	Message    string        `gorm:"column:message" json:"message,omitempty"`
}

type BookingResponse struct {
	ID         uuid.UUID     `json:"id"`
	CustomerID uint          `json:"customer_id"`
	InvoiceID  uuid.UUID     `gorm:"invoice_id" json:"invoice_id,omitempty"`
	PaymentURL string        `gorm:"payment_url" json:"payment_url,omitempty"`
	EventID    uint          `json:"event_id,omitempty"`
	SeatID     uint          `json:"seat_id,omitempty"`
	Email      string        `gorm:"column:email" json:"email,omitempty"`
	Status     BookingStatus `gorm:"column:status" json:"status,omitempty"`
	Message    string        `gorm:"column:message" json:"message,omitempty"`
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
