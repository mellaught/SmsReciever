package reciever

import (
	"log"

	"github.com/streadway/amqp"
)

func (r *Reciever) runPublisher(nameQ string) {

	ch, err := r.conn.Channel()
	failOnError(err, "Failed to open a channel - Publisher")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		nameQ, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue - Publisher")

	for body := range r.smsChan {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		failOnError(err, "Failed to publish a message - Publisher")
		log.Printf(" [x] Sent %s - Publisher", body)
	}
}
