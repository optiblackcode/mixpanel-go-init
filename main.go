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


// hodor - groww - event 1
	
	// Create and track the first event (Sign Up)
	signupEvent := &mixpanel.Event{
		Name: "Sign Up",
		Properties: map[string]interface{}{
			"token":       token, // Pass the token in the properties
			"$user_id":    "USER_ID",
			"is_cart": false,
			"is_collection": false,
			"is_microsite": false,
			"is_product": true,
			"mixpanel_library": "android",
			"$insert_id"
			"screen_name": "PDP",
			"screen_slug": "just-herbs-body-spray-musk-divine-long-lasting-deodorant-spray-for-men-150-ml-8906107054330-240323",
			"user_channel": "android",
			"user_id": "",
			"user_state": "Logged In",
			// events properties to be added
			
			"Signup Type": "Referral",
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
			// all info of user table in poduction 
		},
	}

	// Use PeopleSet to send the user properties to Mixpanel
	if err := mp.PeopleSet(ctx, []*mixpanel.PeopleProperties{peopleProperties}); err != nil {
		log.Fatalf("Error setting People properties: %v", err)
	}

// hodor - groww - event 2

// The Insert ID should be made up of unique attributes in the event that separate it from other performance data. Using our above event example, the uniquely identifiable properties are:

// The ad network name
// The date of the performance data
// The campaign ID
// If we were to send this data more than once to Mixpanel, we know that these 3 properties will always be constant. We can build an Insert ID from that information:

// "G" = Google Ads
// "2023-04-01" = The date of our data
// "12345" = The specific campaign ID
 
// $insert_id = `G-2023-04-01-12345`;

////Note: Keep in mind the Insert ID length limitations. If your ad network has long campaign IDs or other unique properties to use, you should use MD5 or another hashing algorithm to shorten your Insert ID.

	// Create and track another event
	loginEvent := &mixpanel.Event{
		Name: "registration_step_completed",
		Properties: map[string]interface{}{
			"token":    token, // Pass the token in the properties
			"$user_id": "USER_ID",
			"step_order": "1",
             "step_name":"basic_details",
			 "basic_details": , {}   // {{basic_details}}
			 "$insert_id"
			"action":   "Button Click",
		},
	}

	if err := mp.Track(ctx, []*mixpanel.Event{loginEvent}); err != nil {
		log.Fatalf("Error tracking Logged In event: %v", err)
	}

	fmt.Println("All events and People properties successfully tracked!")
}
