package consumer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/mellaught/SmsReciever/src/app/database"

	"github.com/mellaught/SmsReciever/src/app/models"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Consumer struct {
	DB    *db.DataBase
	conn  *amqp.Connection
	queue *amqp.Queue
	ch    *amqp.Channel
}

// Init consumer: dbsql - current sql database, conn - current RabbitMQ conntection, queue - queue's name.
func InitConsumer(dbsql *sql.DB, conn *amqp.Connection, queue string) *Consumer {

	c := &Consumer{
		DB:   &db.DataBase{},
		conn: conn,
	}

	// Initialization database
	db, err := db.InitDB(dbsql)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	c.DB = db

	ch, err := c.conn.Channel()
	failOnError(err, "Failed to open a channel - Consumer")
	c.ch = ch

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue - Consumer")

	c.queue = &q

	return c
}

func (c *Consumer) Run() {

	msgs, err := c.ch.Consume(
		c.queue.Name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer - Consumer")

	forever := make(chan bool)
	// Run consumer.
	go func() {
		for d := range msgs {
			sms := &models.SMSReq{}
			err := json.Unmarshal(d.Body, sms)
			if err != nil {
				log.Println("Consumer decode json error: ", err)
			}
			log.Printf(" [x] %v", sms)
			// PUT in DB.
			// IF COMMIT SUCCES -> delete from queene.
			err = c.DB.PutSMS(sms)
			if err != nil {
				fmt.Println("PUT ERROR:", err)
				// Put into queue again.
				d.Nack(false, true)
				continue
			}
			// Delete from queue.
			d.Ack(false)
		}
	}()
	<-forever
}
