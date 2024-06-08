package handlers

import (
	"fmt"
	"net/http"
	"web-app/pkg/config"
	"web-app/pkg/models"
	"web-app/pkg/render"
)

var Repo *Repository

// create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

type Repository struct {
	App *config.AppConfig
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello there school and college"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	fmt.Println(remoteIP)
	stringMap["remote-ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
