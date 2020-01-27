package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mellaught/SmsReciever/src/app/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DataBase struct {
	DB *sql.DB
}

// InitDB creates tables USERs or SALEs if tables not exists.
func InitDB(db *sql.DB) (*DataBase, error) {

	d := DataBase{
		DB: db,
	}

	// Creates table SMS for Users' sms if not exists
	_, err := db.Exec(CREATE_SMS_IF_NOT_EXISTS)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// PutSMS puts new user's bitcoin address.
func (d *DataBase) PutSMS(sms *models.SMSReq) error {
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	commitTx := false
	defer CloseTransaction(tx, &commitTx)

	stmt, err := tx.Prepare("INSERT INTO SMS(phone, text)  VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	res, err := stmt.Exec(sms.Phone, sms.Text)
	if err != nil {
		return err
	}
	fmt.Println(res)

	commitTx = true
	return nil
}

// IF success -> commit and return result.
// ELSE -> rollbaback.
func CloseTransaction(tx *sql.Tx, commit *bool) {
	if *commit {
		log.Println("Commit sql transaction")
		if err := tx.Commit(); err != nil {
			log.Panic(err)
		}
	} else {
		log.Println("Rollback sql transcation")
		if err := tx.Rollback(); err != nil {
			log.Panic(err)
		}
	}
}
