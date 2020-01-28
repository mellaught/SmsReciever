package app

import (
	"database/sql"
	"log"
	"net/http"

	reciever "github.com/mellaught/SmsReciever/src/app/sms-reciever"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type App struct {
	Router *mux.Router
	rcvr   *reciever.Reciever
}

// InitService is initializes the app.
func NewApp(dbsql *sql.DB, conn *amqp.Connection, name string) *App {

	a := App{
		Router: mux.NewRouter(),
		rcvr:   &reciever.Reciever{}, // Create Reciever and run consumer.
	}

	a.Router = mux.NewRouter()
	a.rcvr = reciever.NewReciever(dbsql, conn, name)
	// Set routers
	a.setRouters()

	return &a
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) setRouters() {
	// Routing for handling the put sms for user.
	a.Post("/sms", a.rcvr.PutSMS)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
