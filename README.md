
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

* database --- Description of  database: driver(**default:** "postgres"), name of database, user, user's password.
* amqp --- Description of RabbitMQ: queue_name, user, user's password, host(**default:** "localhost"), port(**default:** "5672").
* server.host --- Host of **Sms Reciever**. Type: *String.*  **Default:** "localhost".
* server.port --- Port of **Sms Reciever**. Type *String.*  **Default:** "8080".
 
<div  id='sms'/>

  

### 1. *POST* /sms

  

#### Description.

UNDER CONSTRUCTION.

#### Request.

- "Content-Type", "application/json"
- `POST http://service.host:service.port/sms`

*Body:*
```
{
	"phone": string // User id
	"text": int //current minimum number(N) of occurrences, N > 1
	"sex": int // 1 - woman, 2 - man
	"Message": bool // can write private message
}
```  

#### Responce .

`1. StatusCode = 200:`

*Body:*
```
{
	"text": "",
	"responce": 
}
```
#### Cases.
1. - "text": **"The list is empty"**. Means that intersection is empty with current params from request.
	- "responce": ""
2. - "text": **"We found N people"**. Means that intersection is not empty and contains N people with current params from request.
	- "responce": [id_1, id_2, ..., id_N]. Contains people's id from intersection. *Type* **[]int64**.

#### Examples.
 Returned group members with ids: **[1, 2, 4, 6, 1, 3, 1, 2, 3]**.
 - IF *intersect_number* = **2** in request ---> intersection is **[1, 2, 3]**.
- ELSE IF *intersect_number* = **3** in request ---> intersection is **[1]**.
- ELSE IF *intersect_number* = **4** in request ---> intersection is **[ ]**.

**Other parameters from the request impose additional constraints.**

## TODO
-  [x] Put sms' into database.
-  [x] **Tests** for datebase methods.
-  [ ] **Other Tests** .
