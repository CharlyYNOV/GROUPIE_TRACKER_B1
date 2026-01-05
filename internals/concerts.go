package internals

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ArtistDetailsHandler(w http.ResponseWriter, r *http.Request) {



	
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(Artist.ID))
	if err == nil {
		defer resp.Body.Close()

		var locData struct {
			Locations []string `json:"locations"`
		}
		json.NewDecoder(resp.Body).Decode(&locData)

		for i, city := range locData.Locations {
			locData.Locations[i] = strings.ReplaceAll(city, "-", ", ")
			locData.Locations[i] = strings.ReplaceAll(locData.Locations[i], "_", " ")
		}
		Artist.Locations = locData.Locations
	}
	tmpl.ExecuteTemplate(w, "artist.html", Artist)
}
