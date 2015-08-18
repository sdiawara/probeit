package template

import (
	"github.com/sdiawara/probeit/models"
	"html/template"
	"log"
	"net/http"
)

type PollsPage struct{ Polls []models.Probe }

func PollTemplateHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.New("poll").ParseFiles("static/polls.html"))

	pageParam := PollsPage{Polls: models.AllProbes()}
	if err := tmpl.Execute(writer, pageParam); err != nil {
		log.Fatalf("Erreur dans le template : %s", err.Error())
	}
}
