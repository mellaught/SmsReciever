
# Sms Reciever.

 
## Description.

  
Test task for Messaggio. 
Stack: Golang, PostgreSQL, RabbitMQ. Without framework for db(task condition).

## List of endpoints.

[1. **POST**  /sms](#sms)
  

## TORUN.

- `go get github.com/mellaught/SmsReciever`.
- cd $GOPATH/src/github.com/mellaught/SmsReciever/src
- `dep ensure`.
- create `config.json`
- `go run main.go`

### Config file.
Example config in `config.json.example`.
```
{
	"name": "SMS reciever",
	"database": {
     "driver": "postgres",
     "name": "",
     "user": "",
     "password": ""
    },
    "amqp": {
      "queue_name": "",
      "user": "",
      "password": "",
      "host": "localhost",
      "port": "5672"
    },
    "server": {
      "host": "localhost",
      "port": "8000"
    }
}
```

#### Default:

* database --- Description of  database: driver(**default:** "postgres"), name of database, user, user's password.  Type: *String.*
* amqp --- Description of RabbitMQ: queue_name, user, user's password, host(**default:** "localhost"), port(**default:** "5672").  Type: *String.*
* server.host --- Host of **Sms Reciever**. Type: *String.*  **Default:** "localhost".
* server.port --- Port of **Sms Reciever**. Type *String.*  **Default:** "8080".
 
<div  id='sms'/>

  

### 1. *POST* /sms

#### Description.

Put sms into database use RabbitMQ as queue.

#### Request.

- "Content-Type", "application/json"
- `POST http://service.host:service.port/sms`

*Body:*
```
{
	"phone": string // Phone number from sms
	"text": string //Text message from sms
}
```  

#### Responce .

`1. StatusCode = 200:`

*Body:*
```
{
	"text": "Add to queue!"
}
```

## TODO
-  [x] Put sms' into database.
-  [x] **Tests** for datebase methods.
-  [ ] **Other Tests** .
