// Package vendors ..
package vendors

import (
	"fmt"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type NotificationData struct {
	Type string
	ID   string
}

type NotificationMessage struct {
	Body  string
	Title string
	Data  NotificationData
}

func SendNotification(tokens []expo.ExponentPushToken, message NotificationMessage, data NotificationData) {

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       tokens,
			Body:     message.Body,
			Data:     map[string]string{"type": data.Type, "id": data.ID},
			Sound:    "default",
			Title:    message.Title,
			Priority: expo.DefaultPriority,
		},
	)
	// Check errors
	if err != nil {
		fmt.Println(err)
		return
	}
	// Validate responses
	if response.ValidateResponse() != nil {
		fmt.Println(response.PushMessage.To, "failed")
	}
}
