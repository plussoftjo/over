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
type JsonDataType struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
type NotificationMessage struct {
	Body  string
	Title string
	Data  JsonDataType
}

func SendNotification(tokens []expo.ExponentPushToken, message NotificationMessage) {

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       tokens,
			Body:     message.Body,
			Data:     map[string]string{"type": message.Data.Type, "data": message.Data.Data},
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
