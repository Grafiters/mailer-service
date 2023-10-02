package main

import (
    "os"
    "log"

	"mailer/rabbitmq"
	"mailer/token"
	"mailer/interfaces"

	"mailer/handlers"
	"mailer/mailers"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

    queueName := "mailer_queue"
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

    for msg := range messages {
        log.Printf("Received message: %s", msg.Body)

		var recordInterfaces interfaces.Record
		record := token.DecodeToken(string(msg.Body), &recordInterfaces)

		if record != nil {
			log.Println(record)
		}else{
			data, err := handlers.ClaimAllRecordData(string(msg.Body), recordInterfaces.Tag)
			if err != nil {
				log.Println(err)
			}

            log.Println("======= Record =======")
            log.Println(data)
			mailers.SendMail(recordInterfaces, data)
		}
    }
}