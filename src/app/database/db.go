package db

import (
	"github.com/mrKitikat/SmsReciever/src/app/models"
	"database/sql"

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
	_, err := d.DB.Exec("INSERT INTO SMS(phone, text)  VALUES ($1, $2)", sms.Phone, sms.Text)
	if err != nil {
		return err
	}

	return nil
}
