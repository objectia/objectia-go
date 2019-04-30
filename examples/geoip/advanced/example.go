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

	options := &objectia.GeoLocationOptions{
		DisplayFields: "country_code", // Return only country code
	}

	result, err := client.GeoLocation.Get("8.8.8.8", options)
	if err != nil {
		log.Fatal("Failed to get IP location")
	}
	fmt.Printf("Location: %+v\n", result)
}
