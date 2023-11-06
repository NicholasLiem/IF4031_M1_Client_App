package main

import (
	"log"
	"net/http"
	"os"

	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter"
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/clients"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/app"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	/**
	Loading .env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	Setting up http client
	*/
	bookingAPIURL := os.Getenv("BASE_TICKET_APP_URL")
	apiIdentifierToken := os.Getenv("CLIENT_API_KEY")
	headers := map[string]string{
		"Authorization": "Bearer " + apiIdentifierToken,
		"Content-Type":  "application/json",
	}
	restClient := clients.NewRestClient(bookingAPIURL, headers)

	/**
	Setting up DB
	*/
	db := repository.SetupDB()

	/**
	Seeder DB
	*/
	// seeder

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO(db)
	userService := service.NewUserService(dao)
	authService := service.NewAuthService(dao)
	bookingService := service.NewBookingService(dao)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		*restClient,
		userService,
		authService,
		bookingService,
	)

	/**
	Run DB Migration
	*/
	datastruct.Migrate(db, &datastruct.User{}, &datastruct.Booking{})

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
