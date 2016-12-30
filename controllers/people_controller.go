package controllers

import (
	"github.com/nitishparkar/jobberknoll/models"
	"net/http"
	"text/template"
)

type PeopleController struct {
	Template *template.Template
}

func (this *PeopleController) Index(w http.ResponseWriter, r *http.Request) {
	people, err := models.FetchPeople()

	if err == nil {
		w.Header().Add("Content-Type", "text/html")
		this.Template.Execute(w, people)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
	}
}
