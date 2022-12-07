package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	// Load the environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialise the required mail message variables
	from := mail.NewEmail(os.Getenv("SEND_FROM_NAME"), os.Getenv("SEND_FROM_ADDRESS"))
	subject := "Let's Send an Email With Golang and SendGrid"
	to := mail.NewEmail(os.Getenv("SEND_TO_NAME"), os.Getenv("SEND_TO_ADDRESS"))
	plainTextContent := "Here is your AMAZING email!"
	htmlContent := "Here is your <strong>AMAZING</strong> email!"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// Attempt to send the email
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Unable to send your email")
		log.Fatal(err)
	}

	// Check if it was sent
	statusCode := response.StatusCode
	if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		fmt.Println("Email sent!")
	}
}
