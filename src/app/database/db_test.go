package db

import (
	"database/sql"
	"math/rand"
	"testing"

	"github.com/mrKitikat/SmsReciever/src/app/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestPutSMS(t *testing.T) {
	dbsql, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		t.Fatal(err)
	}

	db, err := InitDB(dbsql)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		sms := &models.SMSReq{
			Phone: generateRandomString(11),
			Text:  generateRandomString(100),
		}

		err = db.PutSMS(sms)
		if err != nil {
			t.Fatalf("Put error: %v", err)
		}
	}
}
