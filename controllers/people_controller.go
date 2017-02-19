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

		if err == nil && person.Id != 0 {
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

type NewPersonController struct {
	Template *template.Template
}

func (this *NewPersonController) New(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	w.Header().Add("Content-Type", "text/html")
	this.Template.Execute(w, &person)
}

func (this *NewPersonController) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	bio := r.FormValue("bio")

	person, err := models.SavePerson(name, bio)

	if err == nil {
		http.Redirect(w, r, "/people/" + strconv.Itoa(person.Id), 302)
	} else {
		w.Header().Add("Content-Type", "text/html")
		this.Template.Execute(w, person)
	}
}

func (this *NewPersonController) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawId := vars["id"]

	id, err := strconv.Atoi(rawId)

	if err == nil {
		person, err := models.FetchPerson(id)

		if err == nil && person.Id != 0 {
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

func (this *NewPersonController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawTd := vars["id"]

	id, err := strconv.Atoi(rawTd)

	if err == nil {
		person, err := models.FetchPerson(id)

		if err == nil && person.Id != 0 {
			name := r.FormValue("name")
			bio := r.FormValue("bio")

			person, err := models.UpdatePerson(person, name, bio)
			if err == nil {
				http.Redirect(w, r, "/people/" + strconv.Itoa(person.Id), 302)
			} else {
				w.Header().Add("Content-Type", "text/html")
				this.Template.Execute(w, &person)
			}
		} else {
			w.WriteHeader(404)
			w.Write([]byte(http.StatusText(404)))
		}

	} else {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}
}
