package controler

import (
	"groupie_tracker/internals"
	"html/template"
	"log"
	"net/http"
)

func Concerts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		w.WriteHeader(http.StatusOK)
		return
	}

	tmpl, err := template.ParseFiles("./templates/concerts.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Error while loading page", http.StatusInternalServerError)
		return
	}
	data := struct {
		Artists   []internals.Artist
		Locations []internals.Location
	}{
		Artists:   internals.Artists,
		Locations: internals.Locations,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}
