package main

import (
	"fmt"
	"log"
	"os"

	"github.com/objectia/objectia-go"
)

func main() {
	apiKey := os.Getenv("OBJECTIA_APIKEY")
	client, err := objectia.NewClient(apiKey, nil)

	message := objectia.NewMessage("from@example.com", "Test", "This is just a test", "to@example.com")
	receipt, err := client.Mail.Send(message)
	if err != nil {
		log.Fatal("Failed to send mail")
	}
	fmt.Printf("Location: %+v\n", receipt)
}
