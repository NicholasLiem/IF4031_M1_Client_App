package dto

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
