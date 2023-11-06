package seeder

import (
	"log"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
	// Seed Users
	seedUsers(db)
}

func seedUsers(db *gorm.DB) {
	admin := datastruct.User{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@example.com",
		Password:  "admin",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	admin.Password = string(hashedPassword)

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to seed user: %v", err)
	}
}
