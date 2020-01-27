package reciever

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/mellaught/SmsReciever/src/app/database"
	"github.com/mellaught/SmsReciever/src/app/handler"

	"github.com/streadway/amqp"
)

type Reciever struct {
	smsChan   chan []byte
	queueName string
	conn      *amqp.Connection
	DB        *db.DataBase
}

func NewReciever(conn *amqp.Connection, name string, dbsql *sql.DB) *Reciever {
	r := Reciever{
		smsChan:   make(chan []byte, 0),
		queueName: name,
		conn:      conn,
		DB:        &db.DataBase{},
	}

	db, err := db.InitDB(dbsql)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	r.DB = db

	return &r
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
	// Decode request
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)
	smsBody := buffer.Bytes()
	// Put new sms into channel
	c.smsChan <- smsBody
	handler.ResponJSON(w, http.StatusOK, "Add to queue!")
	return
}
