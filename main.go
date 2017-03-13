package main

import (
	"github.com/gorilla/mux"
	"github.com/nitishparkar/jobberknoll/controllers"
	"github.com/nitishparkar/jobberknoll/models"
	"net/http"
	"os"
	"text/template"
)

var port = ":9090"

func init() {
	err := models.ConnectToDb()
	if err != nil {
		panic(err)
	}

	models.MigrateDb()
}

func main() {
	defer models.CloseDbConnection()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	templates := populateTemplates()

	pc := new(controllers.PeopleController)
	pc.Template = templates.Lookup("index.html")
	router.HandleFunc("/people", pc.Index)

	npc := new(controllers.NewPersonController)
	npc.Template = templates.Lookup("new.html")
	router.HandleFunc("/people/new", npc.New)
	router.HandleFunc("/people/create", npc.Create)

	personController := new(controllers.PersonController)
	personController.Template = templates.Lookup("show.html")
	router.HandleFunc("/people/{id}", personController.Show)

	epc := new(controllers.NewPersonController)
	epc.Template = templates.Lookup("edit.html")
	router.HandleFunc("/people/{id}/edit", epc.Edit)
	router.HandleFunc("/people/{id}/update", epc.Update)

	interactionsController := new(controllers.NewInteractionController)
	interactionsController.Template = templates.Lookup("new_interaction.html")
	router.HandleFunc("/people/{personId}/interactions/new", interactionsController.New)
	router.HandleFunc("/people/{personId}/interactions/create", interactionsController.Create)


	http.Handle("/", router)

	println("Starting web server at port", port)
	http.ListenAndServe(port, nil)
	println("Web server stopped")
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

	_, err := result.ParseFiles(*templatePaths...)

	if err != nil {
		panic(err)
	}

	return result
}
