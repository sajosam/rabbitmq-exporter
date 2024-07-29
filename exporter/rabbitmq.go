package exporter

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (s *Exporter) RabbitMQExport(body []interface{}, routingKey string) {
	// port := os.Getenv("port")
	// host := os.Getenv("host")
	// usrname := os.Getenv("username")
	// passwd := os.Getenv("password")
	// // conn_string := fmt.Sprintf("amqp://%s:%s@%s:%s", usrname, passwd, host, port)
	// pp.Println(conn_string)
	conn, err := amqp.Dial("amqp://onepane:onepane@139.59.58.203:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"op-org", // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bodyBytes, err := json.Marshal(body)
	failOnError(err, "Failed to marshal body to JSON")

	err = ch.PublishWithContext(ctx,
		"op-org",   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s: %s", routingKey, bodyBytes)
}
