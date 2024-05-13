package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"todo/internal/tasker"
)

var templates *template.Template

func SetMux(port string) error {
	var err error
	templatesDir := path.Join("storage", "templates")
	templates, err = template.ParseFiles(filepath.Join(templatesDir, "pages", "tasks.html"), filepath.Join(templatesDir, "blocks", "blocks.html"))
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", listTasks)

	fmt.Println("\nStarting server on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Printf("Error starting server:\n%v", err)
		return err
	}
	return nil
}

func listTasks(w http.ResponseWriter, _ *http.Request) {
	tmpl := templates.Lookup("tasks")
	if tmpl == nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	err := tmpl.Execute(w, tasker.Tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
