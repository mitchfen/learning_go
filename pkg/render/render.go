package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Cache is a map
// We access items in the map via strings, and items in the map are pointers
// because that's what template.ParseFiles returns
var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	var templatePointer *template.Template

	_, inCache := templateCache[templateName]

	// If template is not in the cache, add it
	if !inCache {
		err := updateTemplateCache(templateName)
		if err != nil {
			fmt.Println(err)
		}
	}

	templatePointer = templateCache[templateName]
	log.Println(templateCache)
	err := templatePointer.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// This is a local function and cannot be used outside this package
func updateTemplateCache(templateName string) error {
	layoutTemplates, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		log.Println(err)
	}
	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
	}
	templates = append(templates, layoutTemplates...)

	parsedTemplates, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[templateName] = parsedTemplates
	return nil
}
