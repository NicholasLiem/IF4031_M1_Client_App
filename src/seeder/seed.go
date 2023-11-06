package seeder

import (
	"log"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Seed Users
	SeedUsers(db)
}

func SeedUsers(db *gorm.DB) {
	// Generate and seed fake user data
	users := make([]datastruct.User, 0, 2)

	// Admin role
	admin := datastruct.User{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@example.com",
		Password:  "admin",
		Role:      "admin",
	}
	HashPassword(&admin)

	// User role
	user := datastruct.User{
		FirstName: "user",
		LastName:  "user",
		Email:     "user@example.com",
		Password:  "user",
		Role:      "user",
	}
	HashPassword(&user)

	users = append(users, admin)
	users = append(users, user)

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed user: %v", err)
	}
}

func HashPassword(user *datastruct.User) {
	hashedPasswordUser, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	user.Password = string(hashedPasswordUser)
}
