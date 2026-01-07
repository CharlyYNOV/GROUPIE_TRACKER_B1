package controler

import (
	"groupie_tracker/internals"
	"html/template"
	"log"
	"net/http"
)

// Page d'accueil
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		w.WriteHeader(http.StatusOK)
		return
	}

	tmpl, err := template.ParseFiles("./templates/accueil.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Error while loading page", http.StatusInternalServerError)
		return
	}

	displayArtists := internals.Artists
	if len(displayArtists) > 10 {
		displayArtists = displayArtists[:10]
	}

	displayLocations := internals.Locations
	if len(displayArtists) > 3 {
		displayLocations = displayLocations[:3]
	}

	var displayArtistsWithLocs []ArtistWithLocations
	for _, a := range displayArtists {
		displayArtistsWithLocs = append(displayArtistsWithLocs, ArtistWithLocations{
			Artist:        a,
			LocationsList: internals.GetArtistLocations(a.Id),
		})
	}

	data := struct {
		Artists        []internals.Artist
		DisplayArtists []ArtistWithLocations
		Locations      []internals.Location
	}{
		Artists:        internals.Artists,
		DisplayArtists: displayArtistsWithLocs,
		Locations:      displayLocations,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}
