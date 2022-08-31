package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	//    "google.golang.org/api/option"
)

// func getDecodedFireBaseKey() ([]byte, error) {
// 	fireBaseAuthKey := os.Getenv("FIREBASE_AUTH_KEY")

// 	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return decodedKey, nil
// }

func sendPushNotificationToSingleDevice(fcmClient *messaging.Client, title, body, deviceToken string) (string, error) {
	response, err := fcmClient.Send(context.Background(), &messaging.Message{

		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: deviceToken, // it's a single device token
	})

	if err != nil {
		return "", err
	}
	return response, nil
}

func sendPushNotificationToMultipleDevice(fcmClient *messaging.Client, title string, body string, deviceTokens []string) (*messaging.BatchResponse, error) {
	response, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Tokens: deviceTokens, // it's an array of device tokens
	})

	if err != nil {
		return nil, err
	}

	log.Println("Response success count : ", response.SuccessCount)
	log.Println("Response failure count : ", response.FailureCount)
	return response, nil
}
func initialisFireBaseApp() error {
	// decodedKey, err := getDecodedFireBaseKey()
	// if err != nil {
	// 	return err
	// }

	// opts := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}
	opts := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opts)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	// app, err := firebase.NewApp(context.Background(), nil, opts...)
	if err != nil {
		log.Println("Error in initializing firebase app: %s", err)
		return err
	}

	//initialize client
	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		return err
	}

	resp, err := sendPushNotificationToMultipleDevice(fcmClient, "titl", "body", []string{"token1", "token2", "token3"})
	if err != nil {
		return err
	}

	log.Println(resp)
	return nil
}

func main() {

}
