package publisher

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Publisher struct {
	conn  *amqp.Connection
	queue *amqp.Queue
	ch    *amqp.Channel
}

// Init publisher: conn and queue name.
func InitPublisher(conn *amqp.Connection, queue string) *Publisher {

	p := &Publisher{conn: conn}

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel - Publisher")
	p.ch = ch
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue - Publisher")
	p.queue = &q

	return p
}

// Push puts sms into queue
func (p *Publisher) Push(sms []byte) {
	err := p.ch.Publish(
		"",           // exchange
		p.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        sms,
		})
	failOnError(err, "Failed to publish a message - Publisher")
	log.Printf("Sent %s", sms)
}
