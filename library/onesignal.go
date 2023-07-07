package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Notification struct {
	AppID            string            `json:"app_id"`
	Contents         map[string]string `json:"contents"`
	Headings         map[string]string `json:"headings"`
	IncludedSegments []string          `json:"included_segments"`
}

func PushNotif() {
	// Create a notification object with the required data
	notification := Notification{
		AppID: "YOUR_APP_ID",
		Contents: map[string]string{
			"en": "Hello, this is a test notification!",
		},
		Headings: map[string]string{
			"en": "Test Notification",
		},
		IncludedSegments: []string{"All"},
	}

	// Convert the notification object to JSON
	jsonPayload, err := json.Marshal(notification)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Make a POST request to OneSignal API
	resp, err := http.Post(
		"https://onesignal.com/api/v1/notifications",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Display the response
	fmt.Println("Response:", string(respBody))
}
