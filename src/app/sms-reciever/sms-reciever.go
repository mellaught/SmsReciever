package reciever

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/mellaught/SmsReciever/src/app/sms-reciever/consumer"
	"github.com/mellaught/SmsReciever/src/app/sms-reciever/publisher"

	"github.com/mellaught/SmsReciever/src/app/models"

	"github.com/mellaught/SmsReciever/src/app/handler"

	"github.com/streadway/amqp"
)

type Reciever struct {
	publisher *publisher.Publisher
	consumer  *consumer.Consumer
}

// Create new Reciever App: init publisher, init and run consumer with current input params.
// dbsql - current sql database, conn - current RabbitMQ conntection, queue - queue's name.
func NewReciever(dbsql *sql.DB, conn *amqp.Connection, queue string) *Reciever {

	r := Reciever{
		publisher: &publisher.Publisher{},
		consumer:  &consumer.Consumer{},
	}

	// Initialization publisher: create queue.
	r.publisher = publisher.InitPublisher(conn, queue)
	// Initialization consumer: create queue, create consumer.
	r.consumer = consumer.InitConsumer(dbsql, conn, queue)
	// Run consumer.
	go r.consumer.Run()

	return &r
}

// Send request to worker.
func (c *Reciever) PutSMS(w http.ResponseWriter, r *http.Request) {
	// Decode request for check: telephone number && text lenght.
	sms := &models.SMSReq{}
	resp := models.SMSResp{}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)
	smsBytes := buffer.Bytes()
	err := json.Unmarshal(smsBytes, sms)
	if err != nil {
		resp.Text = err.Error()
		handler.ResponJSON(w, http.StatusBadRequest, resp)
		return
	}
	defer r.Body.Close()

	// Check sms.
	// If OK -> queue
	// ELSE -> responce bad request, please check your sms.
	if !checkPhone(sms.Phone) {
		resp.Text = "Please, check your telephone number."
		handler.ResponJSON(w, http.StatusBadRequest, resp)
		return
	}
	if !checkMessage(sms.Text) {
		resp.Text = "Please, check your sms text. Max lenght equals 255 symbols."
		handler.ResponJSON(w, http.StatusBadRequest, resp)
		return
	}

	// Put new sms into queue
	c.publisher.Push(smsBytes)
	// We have put into queue, but has checked all params.
	resp.Text = "Add to database!"

	handler.ResponJSON(w, http.StatusOK, resp)
	return
}

// Check sms: phone number and message text.
// Lenght and corrent phone with regular expression.
func checkMessage(text string) bool {
	if len(text) > 255 || len(text) == 0 {
		fmt.Println(len(text))
		return true
	}

	return true
}

// Check phone number with regular expression.
func checkPhone(number string) bool {
	matchPhone := regexp.MustCompile(`79\d{2}\d{7}`)
	return matchPhone.Match([]byte(number))
}
