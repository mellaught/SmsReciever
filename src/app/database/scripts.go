package db

var CREATE_SMS_IF_NOT_EXISTS = `
create table if not exists sms (
	ID SERIAL,
   	PHONE VARCHAR(50),
   	TEXT VARCHAR(255),
   	PRIMARY KEY (id)
);`
