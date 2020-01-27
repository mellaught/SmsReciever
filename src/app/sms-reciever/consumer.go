package reciever

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/mellaught/SmsReciever/src/app/models"
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
		false,     // exclusive
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
			sms := &models.SMSReq{}
			err := json.Unmarshal(d.Body, sms)
			if err != nil {
				log.Println("Consumer decode json error: ", err)
			}
			log.Printf(" [x] %v", sms)
			// PUT in DB:
			err = r.DB.PutSMS(sms)
			// IF COMMIT SUCCES -> delete from queene
			if err != nil {
				fmt.Println("PUT ERROR:", err)
				d.Ack(true)
				continue
			}
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C - Consumer")
	<-forever
}

// Check sms: phone number and message text.
// Lenght and corrent phone with regular expression.
func checkSMS(sms *models.SMSReq) bool {
	matchPhone := regexp.MustCompile(`79\d{2}\d{7}`)
	if len(sms.Text) > 254 {
		fmt.Println(len(sms.Text))
		return true
	}
	if len(sms.Phone) < 11 || len(sms.Phone) > 12 {
		fmt.Println(len(sms.Text))
		return false
	}

	return matchPhone.Match([]byte(sms.Phone))
}
