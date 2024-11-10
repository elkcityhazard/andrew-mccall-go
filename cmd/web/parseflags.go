package main

import (
	"flag"
	"fmt"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
)

func parseFlags(app *config.AppConfig, boolChan chan<- bool) {

	flag.StringVar(&app.SiteTitle, "site_title", "Andrew M McCall - Bits & Bytes", "Pass in your site title here")
	flag.StringVar(&app.Port, "site_port", ":8080", "This is the port the application runs on")
	flag.StringVar(&app.DSN, "DSN", "", "The database connection string")
	flag.BoolVar(&app.IsProduction, "is_production", false, "is the app in production? (true|false)")

	flag.Parse()

	fmt.Println(app.SiteTitle)

	boolChan <- true

}
