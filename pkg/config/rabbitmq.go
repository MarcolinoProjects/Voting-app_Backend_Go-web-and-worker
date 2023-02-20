package config

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitConfig struct {
	queue   amqp.Queue
	channel *amqp.Channel
	conn    *amqp.Connection
}

func (r *RabbitConfig) Send(body string, extras ...string) {
	contentType := "text/plain"
	queue := r.queue.Name
	if len(extras) > 0 {
		contentType = extras[0]
		if len(extras) > 1 {
			queue = extras[1]
		}
	}
	err := r.channel.Publish(
		"amq.topic", // exchange
		queue,       // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
func (r *RabbitConfig) initialize() *RabbitConfig {

	conn, err := amqp.Dial(getEnv("RABBIT_HOST", "amqp://rabbitmq:rabbitmq@localhost:5672/"))
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	queue, err := ch.QueueDeclare(
		"voting", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")
	_, err = ch.QueueDeclare(
		"events", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")
	err = ch.QueueBind(
		"voting",    // queue name
		"voting",    // routing key
		"amq.topic", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
	err = ch.QueueBind(
		"events",    // queue name
		"events",    // routing key
		"amq.topic", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
	r.queue = queue
	r.channel = ch
	r.conn = conn
	return r
}
func (r *RabbitConfig) shutdown() {
	err := r.channel.Close()
	failOnError(err, "Failed to close channel with rabbit")
	err = r.conn.Close()
	failOnError(err, "Failed to close connection with rabbit")
}
func (r *RabbitConfig) ListenToQueue(queueName string) <-chan amqp.Delivery {
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")
	return msgs
}
