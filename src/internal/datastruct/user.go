package datastruct

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string    `gorm:"column:first_name" json:"first_name,omitempty"`
	LastName  string    `gorm:"column:last_name" json:"last_name,omitempty"`
	Email     string    `gorm:"uniqueIndex:email" json:"email,omitempty"`
	Password  string    `gorm:"column:password;not null" json:"password,omitempty"`
	Role      Role      `gorm:"column:role;not null;default:user" json:"role,omitempty"`
	Bookings  []Booking `gorm:"foreignKey:CustomerID" json:"-"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      Role   `json:"role"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
