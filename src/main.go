package main

import (
	"github.com/mellaught/SmsReciever/src/app"
	"github.com/mellaught/SmsReciever/src/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/streadway/amqp"
)

func main() {

	cfg := config.NewViperConfig()
	// Read config.json
	serverConfig := cfg.ReadServerConfig()
	dbConfig := cfg.ReadDBConfig()
	amqpConfig := cfg.ReadAMQPConfig()

	// Set DataBase 
	
	dataSourseName := "user=" + dbConfig.DBUser + " " + "DBname=" + dbConfig.DBName + " " + "password=" + dbConfig.DBPassword + " " + "sslmode=disable"
	db, err := sql.Open(dbConfig.DBDriver, dataSourseName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set connection to server
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", amqpConfig.User, amqpConfig.Password, amqpConfig.Host, amqpConfig.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create App with default params from config.json
	app := app.NewApp(db, conn, amqpConfig.QueueName)
	fmt.Println("App has started.")
	app.Run(serverConfig.ServerHost + ":" + serverConfig.ServerPort)
}
