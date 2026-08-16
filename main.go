// https://courses.calhoun.io/lessons/les_wdv2_user_model
package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rianeiromiron/lenslocked/controllers"
	"github.com/rianeiromiron/lenslocked/templates"
	"github.com/rianeiromiron/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	var usersC controllers.Users
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
