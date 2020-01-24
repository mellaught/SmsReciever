package reciever

import (
	"TestJunSMS/src/app/handler"
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/streadway/amqp"
)

type Reciever struct {
	smsChan   chan []byte
	queueName string
	conn      *amqp.Connection
}

func NewReciever(conn *amqp.Connection, name string) *Reciever {
	return &Reciever{
		smsChan:   make(chan []byte, 0),
		queueName: name,
		conn:      conn,
	}
}

// Start Reciever app
func (r *Reciever) Run() {
	// Start publisher
	go r.runPublisher(r.queueName)
	time.Sleep(1 * time.Second)
	// Start consumer
	go r.runConsumer(r.queueName)
	// Put in DataBase worker
}

// Send request to worker.
func (c *Reciever) PutSMS(w http.ResponseWriter, r *http.Request) {

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)
	smsBody := buffer.Bytes()

	fmt.Println("Start put")
	// Put new sms into channel
	c.smsChan <- smsBody
	fmt.Println("Put")
	handler.ResponJSON(w, http.StatusOK, "Add to worker!")
	return
}

func (c *Reciever) PutInDB() {

}
