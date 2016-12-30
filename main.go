package main

import (
	"github.com/gorilla/mux"
	"github.com/nitishparkar/jobberknoll/controllers"
	"net/http"
	"os"
	"text/template"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	templates := populateTemplates()

	pc := new(controllers.PeopleController)
	pc.Template = templates.Lookup("index.html")
	router.HandleFunc("/people", pc.Index)

	personController := new(controllers.PersonController)
	personController.Template = templates.Lookup("show.html")
	router.HandleFunc("/people/{id}", personController.Show)

	http.Handle("/", router)
	http.ListenAndServe(":9090", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
