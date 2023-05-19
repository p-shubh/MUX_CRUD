package Controllers

import (
	"CEO_ASSIST/config"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Config struct {
	Collector *config.DataCollector
	*mux.Router
}

func render(w http.ResponseWriter, tmpl string, data interface{}) {
	// TODO remove the hardcoded templates path. Add it to the config.
	layoutPath := filepath.Join("templates", "layout1.gohtml")

	templatePath := filepath.Join("templates", tmpl+".gohtml")

	//http.Redirect(w,r, "/view-norm", http.StatusFound)

	t, err := template.ParseFiles(layoutPath, templatePath)

	if err != nil {
		log.Println("t parsing error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.ExecuteTemplate(w, "layout", data)

	if err != nil {
		log.Println("t executing error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
