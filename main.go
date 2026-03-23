package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tmpl := template.Must(
		template.Must(template.ParseFiles("index.html")).ParseGlob("components/*.html"),
	)
	r.Handle("/output.css", http.FileServer(http.Dir(".")))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
