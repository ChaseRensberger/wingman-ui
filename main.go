package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type TeamMember struct {
	Name     string
	Email    string
	Initials string
	Role     string
}

type SharedUser struct {
	Name     string
	Email    string
	Initials string
	Access   string
}

type Payment struct {
	Status string
	Email  string
	Amount string
}

type PageData struct {
	TeamMembers []TeamMember
	SharedUsers []SharedUser
	Payments    []Payment
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tmpl := template.Must(
		template.Must(template.ParseFiles("index.html")).ParseGlob("components/*.html"),
	)

	r.Handle("/output.css", http.FileServer(http.Dir(".")))

	data := PageData{
		TeamMembers: []TeamMember{
			{Name: "Sofia Davis", Email: "m@example.com", Initials: "S", Role: "Owner"},
			{Name: "Jackson Lee", Email: "p@example.com", Initials: "J", Role: "Developer"},
			{Name: "Isabella Nguyen", Email: "i@example.com", Initials: "I", Role: "Billing"},
		},
		SharedUsers: []SharedUser{
			{Name: "Olivia Martin", Email: "m@example.com", Initials: "O", Access: "Can edit"},
			{Name: "Isabella Nguyen", Email: "b@example.com", Initials: "I", Access: "Can view"},
			{Name: "Sofia Davis", Email: "p@example.com", Initials: "S", Access: "Can view"},
		},
		Payments: []Payment{
			{Status: "success", Email: "ken99@example.com", Amount: "$316.00"},
			{Status: "success", Email: "abe45@example.com", Amount: "$242.00"},
			{Status: "processing", Email: "monserrat44@example.com", Amount: "$837.00"},
			{Status: "failed", Email: "carmella@example.com", Amount: "$721.00"},
			{Status: "pending", Email: "jason78@example.com", Amount: "$450.00"},
			{Status: "success", Email: "sarah23@example.com", Amount: "$1,280.00"},
		},
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
