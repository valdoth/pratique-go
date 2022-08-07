package handlers

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/valdoth/mygoprogram/config"
	"github.com/valdoth/mygoprogram/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["owner"] = "Tsiaro"

	renderTemplate(w, "home", &models.TemplateData{
		StringData: names,
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	foo := make(map[string]int)
	foo["owner"] = 45

	renderTemplate(w, "about", &models.TemplateData{
		IntData: foo,
	})
}

var appConfig *config.Config

func CreateTemplates(app *config.Config ) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	// t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// t.Execute(w, nil)

	templateCache := appConfig.TemplateCache

	tmpl, ok := templateCache[tmplName + ".page.tmpl"]

	if !ok {
		http.Error(w, "Le template n'existe pas", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)

	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if 0 < len(layouts) {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
