package models

//Config for config.json
type ServerConfig struct {
	DB         *DBConfig
	AMQP       *AMQPConfig
	ServerHost string
	ServerPort string
}

type DBConfig struct {
	DBDriver   string // DataBase driver
	DBName     string // DataBase name
	DBUser     string // DataBase's user
	DBPassword string // User's password
}

type AMQPConfig struct {
	QueueName string // AMQP queue name
	User      string // AMQP user name
	Password  string // AMQP user's password
	Host      string // AMQP host
	Port      string // AMQP port
}

// Request for API method /sms
type SMSReq struct {
	Phone string `json:"phone"` // telephone number, for instance: +79995341054
	Text  string `json:"text"`  // Simple text
}

// Responce for API method /sms
type SMSResp struct {
	Text string `json:"text"` // Responce message.
}
