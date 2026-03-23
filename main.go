package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type pageData struct {
	TestValue int
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tmpl := template.Must(template.ParseFiles("index.html"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, pageData{
			TestValue: 1,
		})
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
