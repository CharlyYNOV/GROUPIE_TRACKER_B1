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

	// On appelle la fonction de internals pour avoir le JSON de la map
	markers := internals.GetMarkersJSON()

	// On prépare les données pour le template
	data := struct {
		Artists     []internals.Artist
		Locations   []internals.Location
		MarkersJSON template.HTML
	}{
		Artists:     internals.Artists,
		Locations:   internals.Locations,
		MarkersJSON: markers,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}
