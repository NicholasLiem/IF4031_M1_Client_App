package datastruct

import "gorm.io/gorm"

type UserModel struct {
	UserID    uint       `gorm:"uniqueIndex:user_id;autoIncrement" json:"user_id,omitempty"`
	FirstName string     `gorm:"column:first_name" json:"first_name,omitempty"`
	LastName  string     `gorm:"column:last_name" json:"last_name,omitempty"`
	Email     string     `gorm:"uniqueIndex:email" json:"email,omitempty"`
	Password  string     `gorm:"column:password;not null" json:"password,omitempty"`
	Role      Role       `gorm:"column:role;not null;default:'user'" json:"role,omitempty"`
	Documents []Document `gorm:"foreignKey:UserRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents,omitempty"`
	gorm.Model
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
