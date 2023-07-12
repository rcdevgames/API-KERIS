package notification

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func PushNotifV2() {
	// Initialize the Firebase app
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json") // Replace with your own service account key file
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Get a messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Failed to get FCM client: %v", err)
	}

	// Construct the message
	message := &messaging.Message{
		Token: "YOUR_DEVICE_TOKEN",
		Notification: &messaging.Notification{
			Title: "Your notification title",
			Body:  "Your notification body",
		},
		Data: map[string]string{
			"key": "value",
		},
	}

	// Send the message
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalf("Failed to send FCM message: %v", err)
	}

	fmt.Println("Push notification sent successfully!")
	fmt.Printf("Response: %+v\n", response)
}
