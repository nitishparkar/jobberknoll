package controllers

import (
	"text/template"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/nitishparkar/jobberknoll/models"
	"github.com/nitishparkar/jobberknoll/viewmodels"
)

type NewInteractionController struct {
	Template *template.Template
}


func (self *NewInteractionController) New(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawId := vars["personId"]

	id, err := strconv.Atoi(rawId)

	if err == nil {
		person, err := models.FetchPerson(id)

		if err == nil && person.ID != 0 {
			vm := viewmodels.InteractionForm{Person: person, Interaction: models.Interaction{}}
			w.Header().Add("Content-Type", "text/html")
			self.Template.Execute(w, &vm)
		} else {
			w.WriteHeader(404)
			w.Write([]byte(http.StatusText(404)))
		}
	} else {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}
}

func (self *NewInteractionController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawId := vars["personId"]

	id, err := strconv.Atoi(rawId)

	if err == nil {
		person, err := models.FetchPerson(id)

		if err == nil && person.ID != 0 {
			intDate := r.FormValue("date")
			intType := r.FormValue("type")
			intDetails := r.FormValue("details")

			interaction, err := models.SaveInteraction(person.ID, intDate, intType, intDetails)

			if err == nil {
				http.Redirect(w, r, "/people/" + strconv.FormatUint(uint64(person.ID), 10), 302)
			} else {
				vm := viewmodels.InteractionForm{Person: person, Interaction: interaction}
				w.Header().Add("Content-Type", "text/html")
				self.Template.Execute(w, &vm)
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

func (self *NewInteractionController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawId := vars["id"]

	id, err := strconv.Atoi(rawId)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	models.DeleteInteraction(id)

	http.Redirect(w, r, "/people/" + vars["personId"], 302)
}