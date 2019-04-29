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

	result, err := client.GeoLocation.GetCurrent()
	if err != nil {
		log.Fatal("Failed to get IP location")
	}
	fmt.Printf("Location: %+v\n", result)
}
