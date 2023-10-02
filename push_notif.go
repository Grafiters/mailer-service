package main

import (
    // "os"
    "log"

	"mailer/rabbitmq"
    "mailer/notification"
    "mailer/token"
    "mailer/interfaces"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

    queueName := "notification_queue"
    conn, err := rabbitmq.ConnectToRabbitMQ()
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatal(err)
    }
    defer ch.Close()

    err = rabbitmq.DeclareQueue(ch, queueName)
    if err != nil {
        log.Fatal(err)
    }

    messages, err := rabbitmq.ConsumeMessages(ch, queueName)
    if err != nil {
        log.Fatal(err)
    }

	log.Printf("[*] Waiting for messages.")

    var recordInterfaces interfaces.Notification
    for msg := range messages {
        log.Printf("Received message: %s", msg.Body)

		record := token.DecodeToken(string(msg.Body), &recordInterfaces)

        log.Println("======= Record =======")
        log.Println(record)
        notification.SendNotification(recordInterfaces, record)
    }
}