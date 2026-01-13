package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ZOLUXERO/go-test/cmd/pkg/config"
	"github.com/ZOLUXERO/go-test/cmd/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template packages
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, templ string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Create template cache, get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	t, ok := tc[templ]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// Render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all of the files name *.page.templ from ./templates
	pages, err := filepath.Glob("./templates/*.page.templ")
	if err != nil {
		return myCache, err
	}

	// Range through all files ending with *.page.templ
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.templ")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.templ")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
