package db

import (
	"database/sql"
	"math/rand"
	"testing"

	"github.com/mellaught/SmsReciever/src/app/models"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	first   = []rune("+7")
	numbers = []rune("100123456789")
)

func generateRandomMessage(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateRandomPhone() string {
	b := make([]rune, 12)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	b[0], b[1] = first[0], first[1]

	return string(b)
}

// Test for DataBase method PutSMS.
// Result: Success: Tests passed.
func TestPutSMS(t *testing.T) {
	dbsql, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		t.Fatal(err)
	}

	db, err := InitDB(dbsql)
	if err != nil || dbsql == nil {
		t.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		sms := &models.SMSReq{
			Phone: generateRandomPhone(),
			Text:  generateRandomMessage(100),
		}

		if i == 5 {
			sms := &models.SMSReq{
				Phone: generateRandomPhone(),
				Text:  generateRandomMessage(1000),
			}
			err = db.PutSMS(sms)
			if err == nil {
				t.Fatalf("Put error: %v", err)
			}
		}

		err = db.PutSMS(sms)
		if err != nil {
			t.Fatalf("Put error: %v", err)
		}
	}
}
