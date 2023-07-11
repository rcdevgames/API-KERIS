package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// FirebasePushNotification represents the data structure for a push notification
type FirebasePushNotification struct {
	To           string               `json:"to"`
	Data         map[string]string    `json:"data"`
	Notification FirebaseNotification `json:"notification"`
}

// FirebaseNotification represents the data structure for a notification message
type FirebaseNotification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Send a push notification
// err := sendPushNotification("your_device_token", "Your notification title", "Your notification body", map[string]string{"key": "value"})
// if err != nil {
// 	fmt.Printf("Failed to send push notification: %v", err)
// 	return
// }

func SendPushNotification(deviceToken, title, body string, data map[string]string) error {
	// Create the push notification object
	pushNotification := FirebasePushNotification{
		To:   deviceToken,
		Data: data,
		Notification: FirebaseNotification{
			Title: title,
			Body:  body,
		},
	}

	// Convert the push notification object to JSON
	payload, err := json.Marshal(pushNotification)
	if err != nil {
		return err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewReader(payload))
	if err != nil {
		return err
	}

	// Set the Firebase server key as the authorization header
	req.Header.Set("Authorization", "key=YOUR_SERVER_KEY")
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("push notification failed with status code %d", resp.StatusCode)
	}

	return nil
}
