package main

import (
	"brankasv1/api/v1/router"
	dbconfig "brankasv1/config/db"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
)

func main() {

	// Print the Welcome Message at Server Startup
	myFigure := figure.NewFigure("BRANKAS", "", true)
	myFigure.Print()
	fmt.Println("")
	slog.Info("Starting Brankas Server on port 4000 ....")

	// Load the environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Setting Up Router
	router := router.Router()

	// Initialize the database connection
	err = dbconfig.InitDB()
	if err != nil {
		fmt.Println("Error initializing database: ", err)
		return
	}
	defer dbconfig.CloseDB() // Close the Database connection

	// Start server
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
