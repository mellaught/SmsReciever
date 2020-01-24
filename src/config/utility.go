package config

import (
	"github.com/mrKitikat/SmsReciever/src/app/models"
)

// Reads AMQP params from config.json
func (v *viperConfig) ReadAMQPConfig() *models.AMQPConfig {
	return &models.AMQPConfig{
		QueueName: v.GetString("amqp.queue_name"),
		User:      v.GetString("amqp.user"),
		Password:  v.GetString("amqp.password"),
		Host:      v.GetString("amqp.host"),
		Port:      v.GetString("amqp.port"),
	}
}

// Reads DataBase params from config.json
func (v *viperConfig) ReadDBConfig() *models.DBConfig {
	return &models.DBConfig{
		DBDriver:   v.GetString("database.driver"),
		DBName:     v.GetString("database.name"),
		DBUser:     v.GetString("database.user"),
		DBPassword: v.GetString("database.password"),
	}
}

// Reads Server params from config.json
func (v *viperConfig) ReadServerConfig() *models.ServerConfig {
	return &models.ServerConfig{
		ServerHost: v.GetString("server.host"),
		ServerPort: v.GetString("server.port"),
	}
}
