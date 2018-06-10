package provider

import (
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
	"github.com/emipochettino/items-api-go/entities"
)

type IQueueProvider interface {
	NotifyItem(item *entities.Item)
}

type QueueProvider struct {
	Url       string
	QueueName string
}

type message struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func (queueProvider QueueProvider) NotifyItem(item *entities.Item) {
	conn, ch, q := queueProvider.getConnection()
	defer func() {
		conn.Close()
		ch.Close()
	}()

	message := message{ID: item.ID, Type: "item"}

	body, _ := json.Marshal(message)

	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	failOnError(err, "Failed to publish a message")
	log.Printf("Message with body %s sent", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (queueProvider *QueueProvider) getConnection() (conn *amqp.Connection, ch *amqp.Channel, q amqp.Queue) {
	//TODO if somethings went wrong here the error stops the app
	conn, err := amqp.Dial(queueProvider.Url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err = ch.QueueDeclare(
		queueProvider.QueueName, // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, q
}
