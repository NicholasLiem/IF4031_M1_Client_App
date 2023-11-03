package datastruct

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	CustomerID uint          `gorm:"customer_id" json:"customer_id,omitempty"`
	EventID    uint          `gorm:"column:event_id" json:"event_id,omitempty"`
	SeatID     uint          `gorm:"column:seat_id" json:"seat_id,omitempty"`
	Status     BookingStatus `gorm:"column:status" json:"status,omitempty"`
	Message    string        `gorm:"column:message" json:"message,omitempty"`
	Customer   UserModel     `gorm:"foreignKey:CustomerID" json:"-"`
}

type BookingStatus string

const (
	BookingFailed    BookingStatus = "failed"
	BookingSuccess   BookingStatus = "success"
	BookingOnProcess BookingStatus = "on-process"
)

func IsValidStatus(status BookingStatus) bool {
	return status == BookingFailed || status == BookingSuccess || status == BookingOnProcess
}
