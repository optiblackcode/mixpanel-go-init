package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mixpanel/mixpanel-go"
)

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the Mixpanel token from the environment
	token := os.Getenv("MIXPANEL_TOKEN")
	if token == "" {
		log.Fatalf("MIXPANEL_TOKEN is not set in the environment")
	}

	// Create a context
	ctx := context.Background()

	// Initialize Mixpanel client
	mp := mixpanel.NewApiClient(token)

	// Create an event
	event := &mixpanel.Event{
		Name: "people_event",
		Properties: map[string]interface{}{
			"token":    token, // Pass the token in the properties
			"$user_id": "12345",
			"action":   "logged_in",
		},
	}

	// Track the event
	if err := mp.Track(ctx, []*mixpanel.Event{event}); err != nil {
		log.Fatalf("Error tracking event: %v", err)
	}

	// Set People Properties
	peopleProperties := &mixpanel.PeopleProperties{
		DistinctID: "12345", // Unique identifier for the user
		Properties: map[string]interface{}{
			"$name":  "Jane Doe",
			"$email": "jane.doe@example.com",
			"plan":   "Premium", // Custom property
		},
	}

	// Use PeopleSet to send the user properties to Mixpanel
	if err := mp.PeopleSet(ctx, []*mixpanel.PeopleProperties{peopleProperties}); err != nil {
		log.Fatalf("Error setting People properties: %v", err)
	}

	fmt.Println("Event and People properties successfully tracked!")
}
