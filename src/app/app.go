package app

import (
	reciever "TestJunSMS/src/app/sms-reciever"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type App struct {
	Router *mux.Router
	Rcvr   *reciever.Reciever
}

// InitService is initializes the app.
func NewApp(dbsql *sql.DB, conn *amqp.Connection, name string) *App {

	a := App{
		Router: mux.NewRouter(),
		Rcvr:   reciever.NewReciever(conn, name),
	}

	a.Router = mux.NewRouter()
	// Start reciever
	a.Rcvr.Run()
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
	a.Post("/sms", a.Rcvr.PutSMS)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
