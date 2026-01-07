package controler

import (
	"groupie_tracker/internals"
	"html/template"
	"log"
	"net/http"
)

type ArtistWithLocations struct {
	internals.Artist
	LocationsList []string
}

// Nouvelle route pour la page des Artistes
func ViewAllArtistsPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		w.WriteHeader(http.StatusOK)
		return
	}

	tmpl, err := template.ParseFiles("./templates/viewAllArtists.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Error while loading page", http.StatusInternalServerError)
		return
	}

	search_input := r.URL.Query().Get("search")
	// searchQuery := search_input

	artists := internals.Artists
	if search_input != "" {
		artists, _ = internals.FilterArtists(artists, search_input)
	}

	// Récupérer les suggestions
	// suggestions := internals.GetSearchSuggestions(searchQuery)

	var artistsWithLocs []ArtistWithLocations
	for _, a := range artists {
		artistsWithLocs = append(artistsWithLocs, ArtistWithLocations{
			Artist:        a,
			LocationsList: internals.GetArtistLocations(a.Id),
		})
	}

	data := struct {
		Artists     []ArtistWithLocations
		Suggestions []string
		SearchQuery string
	}{
		Artists:     artistsWithLocs,
		Suggestions: internals.GetSearchSuggestions(search_input), // garde une seule variable
		SearchQuery: search_input,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}
