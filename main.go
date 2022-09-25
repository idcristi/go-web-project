package main

import (
	"fmt"
	"go-web-dev/controllers"
	"go-web-dev/views"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
