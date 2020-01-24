package reciever

import (
	"github.com/mrKitikat/SmsReciever/src/app/models"
	"encoding/json"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Run consumer
func (r *Reciever) runConsumer(nameQueue string) {

	ch, err := r.conn.Channel()
	failOnError(err, "Failed to open a channel - Consumer")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		nameQueue, // name
		false,     // durable
		false,     // delete when unused
		false,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue - Consumer")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer - Consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			sms := models.SMSReq{}
			err := json.Unmarshal(d.Body, &sms)
			if err != nil {
				log.Println("Consumer decode json error: ", err)
			}
			log.Printf(" [x] %v", sms)
			d.Ack(false)
			// PUT in DB:
			// Create tx
			// IF COMMIT SUCCES -> delete from queene
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C - Consumer")
	<-forever
}
