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

func ViewAllArtistsPage(w http.ResponseWriter, r *http.Request) {
	// On ignore le POST car SearchBar s'occupe de la redirection vers GET
	if r.Method == "POST" {
		r.ParseForm()
		return
	}

	tmpl, err := template.ParseFiles("./templates/viewAllArtists.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Error while loading page", http.StatusInternalServerError)
		return
	}

	// 1. RÉCUPÉRATION DE L'INPUT (Correction : on lit le paramètre ?search= de l'URL)
	search_input := r.URL.Query().Get("search")

	artists := internals.Artists
	found := true

	// 2. FILTRAGE (Correction : on vérifie si FilterArtists renvoie des résultats)
	if search_input != "" {
		artists, found = internals.FilterArtists(artists, search_input)
	}

	var artistsWithLocs []ArtistWithLocations

	// 3. BOUCLE SÉCURISÉE (Correction : on n'ajoute que si 'found' est vrai et 'artists' non nil)
	if found && artists != nil {
		for _, a := range artists {
			artistsWithLocs = append(artistsWithLocs, ArtistWithLocations{
				Artist:        a,
				LocationsList: internals.GetArtistLocations(a.Id),
			})
		}
	}

	// 4. DONNÉES SYNCHRONISÉES (On utilise search_input pour les suggestions et la query)
	data := struct {
		Artists     []ArtistWithLocations
		Suggestions []string
		SearchQuery string
	}{
		Artists:     artistsWithLocs,
		Suggestions: internals.GetSearchSuggestions(search_input),
		SearchQuery: search_input,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}
