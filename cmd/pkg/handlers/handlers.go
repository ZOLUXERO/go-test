package handlers

import (
	"net/http"

	"github.com/ZOLUXERO/go-test/cmd/pkg/config"
	"github.com/ZOLUXERO/go-test/cmd/pkg/models"
	"github.com/ZOLUXERO/go-test/cmd/pkg/render"
)

// Repository pattern
// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler, this uses receivers
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.templ", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Send data to the template
	render.RenderTemplate(w, "about.page.templ", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is the generals page handler, this uses receivers
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.templ", &models.TemplateData{})
}

// Majors is the majors page handler, this uses receivers
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.templ", &models.TemplateData{})
}

// Availability is the availability page handler, this uses receivers
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.templ", &models.TemplateData{})
}

// Contact is the contact page handler, this uses receivers
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.templ", &models.TemplateData{})
}
