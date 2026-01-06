package internals

import (
	"encoding/json"
	"net/http"
)

// Coordonnées des villes
var CityCoords = map[string][2]float64{
	"london-uk":             {51.505, -0.09},
	"paris-france":          {48.8566, 2.3522},
	"lyon-france":           {45.7640, 4.8357},
	"marseille-france":      {43.2965, 5.3698},
	"lisbon-portugal":       {38.7223, -9.1393},
	"madrid-spain":          {40.4168, -3.7038},
	"barcelona-spain":       {41.3851, 2.1734},
	"berlin-germany":        {52.5200, 13.4050},
	"new-york-usa":          {40.7128, -74.0060},
	"los-angeles-usa":       {34.0522, -118.2437},
	"tokyo-japan":           {35.6895, 139.6917},
	"praia-cape_verde":      {14.9177, -23.5092},
	"amsterdam-netherlands": {52.3676, 4.9041},
}

type ConcertPoint struct {
	CityName string
	Lat      float64
	Long     float64
}

type Locations2 struct {
	Index []struct {
		ID       int      `json:"id"`
		Location []string `json:"Locations2"`
		Dates    string   `json:"dates"`
		DatesURL string   `json:"dates"`
	} `json:"index"`
}

func FetchLocations() (Locations2, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/Locations2")
	if err != nil {
		return Locations2{}, err
	}
	defer resp.Body.Close()

	var locs Locations2
	err = json.NewDecoder(resp.Body).Decode(&locs)
	return locs, err
}
