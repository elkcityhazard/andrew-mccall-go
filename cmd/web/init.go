package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
	"github.com/elkcityhazard/andrew-mccall-go/internal/database"
	"github.com/elkcityhazard/andrew-mccall-go/internal/handlers"
	"github.com/elkcityhazard/andrew-mccall-go/internal/mailer"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
	"github.com/elkcityhazard/andrew-mccall-go/internal/repository/sqldbconn"
	"github.com/elkcityhazard/andrew-mccall-go/internal/templates"
	amrouter "github.com/elkcityhazard/andrew-mccall-go/pkg/am_router"
)

const (
	TIMEOUT_DURATION    = time.Second * 30
	CONN_MAX_LIFETIME   = time.Minute * 3
	CONN_MAX_OPEN_CONNS = 10
	CONN_MAX_IDLE_CONNS = 30
	CONN_MAX_IDLE_TIME  = time.Second * 90
)

// msgListern listens for any message and prints it to the console
func msgListener(msgChan chan string) {
	for {
		select {
		case msg := <-msgChan:
			log.Println(msg)
		}
	}
}

// run starts the application
func run() {
	email := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASSWORD")
	app.MsgChan = make(chan string)
	app.TemplateCache = templates.BuildTemplateCache()
	app.WG = &sync.WaitGroup{}
	app.MU = &sync.Mutex{}
	app.Context = context.Background()

	// dummy mailer setup
	mailDispatcher := mailer.New("localhost", 1025, "username", "password", "web@andrew-mccall.com")
	app.Mailer = mailDispatcher
	go mailDispatcher.ListenForIncomingEmail()
	flagChan := make(chan bool)
	go msgListener(app.MsgChan)
	go parseFlags(&app, flagChan)
	<-flagChan
	db, err := connectToDB(&app)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	newSessionManager(&app, db)
	render.NewRenderer(&app)
	dbServicer := sqldbconn.NewSQLDbConn(&app, db)
	handlerRepo := handlers.NewHandlerRepo(&app, dbServicer)
	handlers.SetHandlerRepo(handlerRepo)

	// pass the handler repo into the templates so we can fetch the author without some extra one off model for an inner join
	templates.SetTemplateSQLDbRepo(dbServicer)

	superAdmin := createSeedUser(email, password)

	_, err = dbServicer.InsertUser(superAdmin)

	if err != nil {
		log.Fatalln(err)
	}
	startServer(&app)
}

func createSeedUser(email, password string) *models.User {

	p := models.Password{}

	p.PlainText = password

	pwHash, err := argon2id.CreateHash(p.PlainText, argon2id.DefaultParams)

	if err != nil {
		log.Fatalln(err)
	}
	p.Hash = pwHash

	u := &models.User{
		Email:    email,
		Username: "elkcityhazard",
		Password: &p,
		Role:     "super_admin",
		IsActive: true,
	}

	return u

}

// startServer starts the server on the specified port
func startServer(app *config.AppConfig) {

	fmt.Printf("Starting server on: %s\n", app.Port)

	rtr := amrouter.NewRouter()

	rtr.PathToStaticDir = "/internal/static"

	rtr.Use(LoadSession)
	rtr.Use(stripTrailingSlash)
	rtr.Use(csrfToken)

	rtr.AddRoute("GET", "/", handlers.Repo.HomeGetHandler, requiresAuth)

	srv := &http.Server{
		Addr: app.Port,
		//Handler:           LoadSession(csrfToken(stripTrailingSlash(routes()))),
		Handler:           rtr,
		IdleTimeout:       TIMEOUT_DURATION,
		WriteTimeout:      TIMEOUT_DURATION,
		ReadTimeout:       TIMEOUT_DURATION,
		ReadHeaderTimeout: TIMEOUT_DURATION,
		MaxHeaderBytes:    2 >> 30,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

// connectToDB watis for initDB to get done and returns the database conn and an error
func connectToDB(app *config.AppConfig) (*sql.DB, error) {
	db := make(chan *sql.DB)
	err := make(chan error)
	done := make(chan bool)
	go initDB(app, db, err, done)
	for {
		select {
		case d := <-db:
			return d, nil
		case e := <-err:
			return nil, e
		case <-done:
			close(db)
			close(err)
			close(done)
		}
	}
}

// initDB initializes a new database connection and sends it to the receiver channel
func initDB(app *config.AppConfig, dbChan chan<- *sql.DB, dbErrorChan chan<- error, dbDoneChan chan<- bool) {
	db, err := database.NewDriver(app)
	if err != nil {
		dbErrorChan <- err
		dbDoneChan <- true
	}
	if err = db.Ping(); err != nil {
		dbErrorChan <- err
		dbDoneChan <- true
	}
	db.SetConnMaxLifetime(CONN_MAX_LIFETIME)
	db.SetMaxOpenConns(CONN_MAX_OPEN_CONNS)
	db.SetMaxIdleConns(CONN_MAX_IDLE_CONNS)
	db.SetConnMaxIdleTime(CONN_MAX_IDLE_TIME)
	dbChan <- db
	dbDoneChan <- true
}
