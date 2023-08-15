package rabbitmq

import (
    "os"
    "log"

    "github.com/joho/godotenv"
    "github.com/streadway/amqp"
)

func ConnectToRabbitMQ()( *amqp.Connection, error ) {
    err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Println("Error connecting to RabbitMQ %s", err)
		return nil, err
	}

	return conn, nil
}

func DeclareQueue(ch *amqp.Channel, queueName string) error {
    _, err := ch.QueueDeclare(
        queueName,
        true,
        false,
        false,
        false,
        nil,
    )
    return err
}

func ConsumeMessages(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
    return ch.Consume(
        queueName, // Queue
        "",     // Consumer
        true,   // Auto-sAck
        false,  // Excluive
        false,  // No-local
        false,  // No-wait
        nil,    // Args
    )
}