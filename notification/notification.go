package notification

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"mailer/interfaces"
	"io/ioutil"
)

type Message struct {
	NotificationBody 	NotificationBody 	`json:"notification"`
	RegistrationIds		[]string       		`json:"registration_ids"`
}

type NotificationBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func SendNotification(record interfaces.Notification, msg interface{}) {
	detail := Message{
		NotificationBody: NotificationBody{
			Title: record.Title,
			Body:  record.Message,
		},
		RegistrationIds: record.User.DeviceToken,
	}
	

	log.Println(detail)
	body, err := json.Marshal(detail)
	if err != nil {
		log.Printf("Failed to marshal message: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to create request: %v\n", err)
		return
	}

	req.Header.Set("Authorization", "key="+os.Getenv("SERVER_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
	} else {
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response status: %s", resp.Status)
		log.Printf("Response body: %s", string(respBody))
		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to send message, status: %s", resp.Status)
		} else {
			log.Println("Successfully sent message")
		}
	}
}