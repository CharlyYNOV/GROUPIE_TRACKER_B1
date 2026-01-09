package internals

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func FilterArtists(artists []Artist, input string) ([]Artist, bool) {
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		return artists, true
	}

	var filtered []Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), input) {
			filtered = append(filtered, artist)
		}
	}

	if len(filtered) == 0 {
		return nil, false
	}
	return filtered, true
}

func GetSearchSuggestions(input string) []string {
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		return []string{}
	}

	var suggestions []string
	seen := make(map[string]bool)

	for _, artist := range Artists {
		if strings.Contains(strings.ToLower(artist.Name), input) {
			if !seen[artist.Name] {
				suggestions = append(suggestions, artist.Name)
				seen[artist.Name] = true
			}
		}
	}

	return suggestions
}

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
	encodedInput := url.QueryEscape(searchInput)
	referer := r.Referer()

	if strings.Contains(referer, "/concerts") {
		if searchInput == "" {
			http.Redirect(w, r, "/concerts", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/concerts?search="+encodedInput, http.StatusSeeOther)
		}
	} else {
		if searchInput == "" {
			http.Redirect(w, r, "/artists", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/artists?search="+encodedInput, http.StatusSeeOther)
		}
	}
}
