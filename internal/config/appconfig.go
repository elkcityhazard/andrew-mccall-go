package config

import (
	"context"
	"html/template"
	"sync"

	"github.com/alexedwards/scs/v2"
	"github.com/elkcityhazard/andrew-mccall-go/internal/mailer"
)

type AppConfig struct {
	SiteTitle      string
	Port           string
	DSN            string
	IsProduction   bool
	TemplateCache  map[string]*template.Template
	SessionManager *scs.SessionManager
	Context        context.Context
	MU             *sync.Mutex
	WG             *sync.WaitGroup
	MsgChan        chan string
	Mailer         mailer.Mailer
}
