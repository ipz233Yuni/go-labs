package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

func render(w http.ResponseWriter, page string, title string) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{Title: title}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "home.html", "Головна")
	})

	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		render(w, "history.html", "Історія")
	})

	http.HandleFunc("/places", func(w http.ResponseWriter, r *http.Request) {
		render(w, "places.html", "Визначні місця")
	})

	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		render(w, "gallery.html", "Галерея")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		render(w, "contact.html", "Контакти")
	})

	log.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
