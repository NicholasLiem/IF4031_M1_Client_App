package repository

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"gorm.io/gorm"
)

type BookingQuery interface {
	CreateBooking(booking datastruct.Booking) (*datastruct.Booking, error)
	UpdateBooking(bookingID uint, booking datastruct.Booking) (*datastruct.Booking, error)
	DeleteBooking(bookingID uint) (*datastruct.Booking, error)
	GetBooking(bookingID uint) (*datastruct.Booking, error)
	GetBookingsFromCustomerID(customerID uint) ([]datastruct.Booking, error)
}

type bookingQuery struct {
	pgdb *gorm.DB
}

func NewBookingQuery(pgdb *gorm.DB) BookingQuery {
	return &bookingQuery{
		pgdb: pgdb,
	}
}

func (bq *bookingQuery) CreateBooking(booking datastruct.Booking) (*datastruct.Booking, error) {
	if err := bq.pgdb.Create(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (bq *bookingQuery) UpdateBooking(bookingID uint, booking datastruct.Booking) (*datastruct.Booking, error) {
	existingBooking := datastruct.Booking{}
	result := bq.pgdb.First(&existingBooking, bookingID)
	if result.Error != nil {
		return nil, result.Error
	}

	existingBooking.CustomerID = booking.CustomerID
	existingBooking.EventID = booking.EventID
	existingBooking.SeatID = booking.SeatID
	existingBooking.Status = booking.Status

	result = bq.pgdb.Save(&existingBooking)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingBooking, nil
}

func (bq *bookingQuery) DeleteBooking(bookingID uint) (*datastruct.Booking, error) {
	existingBooking := datastruct.Booking{}
	result := bq.pgdb.First(&existingBooking, bookingID)
	if result.Error != nil {
		return nil, result.Error
	}

	result = bq.pgdb.Delete(&existingBooking)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingBooking, nil
}

func (bq *bookingQuery) GetBooking(bookingID uint) (*datastruct.Booking, error) {
	booking := datastruct.Booking{}
	result := bq.pgdb.First(&booking, bookingID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &booking, nil
}

func (bq *bookingQuery) GetBookingsFromCustomerID(customerID uint) ([]datastruct.Booking, error) {
	var bookings []datastruct.Booking
	err := bq.pgdb.Where("customer_id = ?", customerID).Find(&bookings).Error
	return bookings, err
}
