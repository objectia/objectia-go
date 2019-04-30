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

	result, err := client.GeoLocation.GetBulk([]string{"1.2.3.4", "8.8.8.8"}, nil)
	if err != nil {
		log.Fatal("Failed to get IP location")
	}

	for _, v := range result {
		fmt.Printf("Location: %+v\n", v)
	}
}
