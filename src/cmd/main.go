package main

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/app"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/service"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {

	/**
	Creating context
	*/
	//ctx := context.Background()

	/**
	Loading .env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	Setting up DB
	*/
	db := repository.SetupDB()

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO(db)
	userService := service.NewUserService(dao)
	authService := service.NewAuthService(dao)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		userService,
		authService,
	)

	/**
	Run DB Migration
	*/
	Migrate(db)

	/**
	Setting up the router
	*/
	serverRouter := adapter.NewRouter(*server)

	/**
	Running the server
	*/
	port := os.Getenv("PORT")
	log.Println("Running the server on port " + port)

	if os.Getenv("ENVIRONMENT") == "DEV" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+port, serverRouter))
	}
	log.Fatal(http.ListenAndServe(":"+port, serverRouter))
}

func Migrate(db *gorm.DB) {
	errPdf := db.AutoMigrate(&datastruct.Document{})
	if errPdf != nil {
		return
	}
	err := db.AutoMigrate(&datastruct.UserModel{})
	if err != nil {
		return
	}
}
