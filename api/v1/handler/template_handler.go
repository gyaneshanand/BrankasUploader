package handler

import (
	"html/template"
	"net/http"
	"os"
)

// Serve Document Upload Form
func ServeForm(w http.ResponseWriter, r *http.Request) {

	formTemplate, err := template.New("form.html").ParseFiles("view/form.html")
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	pageVariables := PageVariables{
		AuthToken: os.Getenv("AUTH_TOKEN"),
	}

	err = formTemplate.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, "Unable to render HTML", http.StatusInternalServerError)
	}

}

type PageVariables struct {
	AuthToken string
}
