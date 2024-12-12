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

	// Create and track the first event (Sign Up)
	signupEvent := &mixpanel.Event{
		Name: "Sign Up",
		Properties: map[string]interface{}{
			"token":       token, // Pass the token in the properties
			"Signup Type": "Referral",
			"$user_id":    "USER_ID",
		},
	}

	if err := mp.Track(ctx, []*mixpanel.Event{signupEvent}); err != nil {
		log.Fatalf("Error tracking Sign Up event: %v", err)
	}

	// Set People Properties
	peopleProperties := &mixpanel.PeopleProperties{
		DistinctID: "USER_ID", // Unique identifier for the user
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

	// Create and track another event
	loginEvent := &mixpanel.Event{
		Name: "Logged In",
		Properties: map[string]interface{}{
			"token":    token, // Pass the token in the properties
			"$user_id": "USER_ID",
			"action":   "Button Click",
		},
	}

	if err := mp.Track(ctx, []*mixpanel.Event{loginEvent}); err != nil {
		log.Fatalf("Error tracking Logged In event: %v", err)
	}

	fmt.Println("All events and People properties successfully tracked!")
}
