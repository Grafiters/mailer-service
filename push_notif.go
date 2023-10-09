package main


import (
    "log"

    "mailer/token"
	"mailer/rabbitmq"
    "mailer/notification"
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

        if record != nil {
			log.Println(record)
		}else{
			log.Println("Data => ",recordInterfaces)
            notification.SendNotification(recordInterfaces, record)
		}
    }
}