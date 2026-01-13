package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ZOLUXERO/go-test/cmd/pkg/config"
	"github.com/ZOLUXERO/go-test/cmd/pkg/handlers"
	"github.com/ZOLUXERO/go-test/cmd/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true when in production
	app.InProduction = false

	// manage sessions, you could use (redis, mysql, postgresql)
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Starting application on port:", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
