package internals

import (
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// Cette fonction redirige vers la page de l'artiste demandé
func SearchBar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Printf("Error while parsing search bar input: %v", err)
		http.Error(w, "Error while parsing search bar input", http.StatusInternalServerError)
		return
	}

	searchInput := strings.TrimSpace(r.PostFormValue("search-bar"))
	if searchInput == "" {
		http.Redirect(w, r, "/artists", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/artists?search="+url.QueryEscape(searchInput), http.StatusSeeOther)
}

// Cette fonction retourne une liste avec tous les artistes comportant le nom [input] dans l'input
// mais aussi tous les groupes contenant le nom [input] dans leur liste de membres
func FilterArtists(artists []Artist, input string) []Artist {
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		return artists
	}

	var filtered []Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), input) {
			filtered = append(filtered, artist)
			continue
		}
	}

	return filtered
}

// GetSearchSuggestions retourne jusqu'à 3 suggestions basées sur l'input de l'utilisateur
func GetSearchSuggestions(input string) []string {
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		return []string{}
	}

	var suggestions []string
	seenSuggestions := make(map[string]bool)

	// Chercher dans les noms d'artistes
	for _, artist := range Artists {
		artistName := artist.Name
		if strings.Contains(strings.ToLower(artistName), input) {
			if !seenSuggestions[artistName] {
				suggestions = append(suggestions, artistName)
				seenSuggestions[artistName] = true
			}
		}
	}

	// Chercher dans les membres
	for _, artist := range Artists {
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), input) {
				suggestion := member + " (member of " + artist.Name + ")"
				if !seenSuggestions[suggestion] {
					suggestions = append(suggestions, suggestion)
					seenSuggestions[suggestion] = true
				}
			}
		}
	}

	// Trier par ordre alphabétique
	sort.Strings(suggestions)

	// Limiter à 3 suggestions
	if len(suggestions) > 3 {
		suggestions = suggestions[:3]
	}

	return suggestions
}

func FilterMember(artist []Artist) {

}