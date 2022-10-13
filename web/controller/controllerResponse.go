package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func ShowView(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

	// Specify the path of the view
	pagePath := filepath.Join("web", "tpl", templateName)

	resultTemplate, err := template.ParseFiles(pagePath)
	if err != nil {
		fmt.Printf("Error creating template instance: %v", err)
		return
	}

	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Printf("An error occurred while fusing data in the template: %v", err)
		return
	}

}
