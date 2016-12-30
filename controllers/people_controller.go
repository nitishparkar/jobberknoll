package controllers

import (
	"github.com/gorilla/mux"
	"github.com/nitishparkar/jobberknoll/models"
	"net/http"
	"strconv"
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

type PersonController struct {
	Template *template.Template
}

func (this *PersonController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawId := vars["id"]

	id, err := strconv.Atoi(rawId)

	if err == nil {
		person, err := models.FetchPerson(id)

		if err == nil && person != (models.Person{}) {
			w.Header().Add("Content-Type", "text/html")
			this.Template.Execute(w, &person)
		} else {
			w.WriteHeader(404)
			w.Write([]byte(http.StatusText(404)))
		}
	} else {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}
}
