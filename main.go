package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

var user = "user"
var pass = "Passw0rd"
var target = "http://192.168.30.129"

func main1() {
	// Create a new client.
	client := NewHTTPClient(target, 30*time.Second)

	// Authenticate.
	ctx := context.Background()
	_, err := client.CheckLogin(ctx, user, pass)
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		return
	}

	// Get metrics.
	metrics, err := client.GetMetrics(ctx)
	if err != nil {
		fmt.Printf("Failed to get metrics: %v\n", err)
		return
	}

	// Display metrics as a table.
	DisplayMetricsTable(os.Stdout, metrics)
}

func main2() {
	// Create a new client with a 30-second timeout.
	client := NewHTTPClient(target, 30*time.Second)

	// Authenticate.
	ctx := context.Background()
	token, err := client.CheckLogin(ctx, user, pass)
	fmt.Printf("token: %s\n", token)
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		return
	}
	fmt.Printf("Authenticated with token: %s\n", token)

	// List collections.
	collections, err := client.ListCollections(ctx, 1)
	if err != nil {
		fmt.Printf("Failed to list collections: %v\n", err)
		return
	}
	for _, c := range collections {
		fmt.Printf("Collection: %s (ID: %s)\n", c.Title, c.ID)
	}

	/*
		// Create an item.
		item := &Item{
			Name:       "New Item",
			Collection: IRI("collection_id"), // Replace with valid collection IRI.
			Quantity:   1,
			Visibility: "public",
		}

		createdItem, err := client.CreateItem(ctx, item)
		if err != nil {
			fmt.Printf("Failed to create item: %v\n", err)
			return
		}
		fmt.Printf("Created item: %s (ID: %s)\n", createdItem.Name, createdItem.ID)
	*/
}

// ptr returns a pointer to a string.
func ptr(s string) *string {
	return &s
}

func main() {
	//main1()
	main2()
}
