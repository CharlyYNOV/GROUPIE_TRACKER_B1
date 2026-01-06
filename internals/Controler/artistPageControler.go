package controler

import (
	"groupie_tracker/internals"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Artist ID required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var artist *internals.Artist
	for i := range internals.Artists {
		if internals.Artists[i].Id == id {
			artist = &internals.Artists[i]
			break
		}
	}

	if artist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
	}

	tmpl, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		log.Printf("Error parsing template : %v", err)
		http.Error(w, "Error while loading page", http.StatusInternalServerError)
		return
	}

	data := struct {
		ArtistPage internals.Artist
		Artists    []internals.Artist
	}{
		ArtistPage: *artist,
		Artists:    internals.Artists,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}
