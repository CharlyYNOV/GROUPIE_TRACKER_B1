package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type location_API_Response_Structure struct {
	Index []Location `json:"index"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type relation_API_Response_Structure struct {
	Index []Relation `json:"index"`
}

type Concert struct {
	ArtistName string
	Location   string
	Date       string
	ArtistId   int
}

var (
	Artists     []Artist
	Locations   []Location
	Relations   []Relation
	AllConcerts []Concert
)

func Main_api() {
	data, err := api("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal("Error fetching API 1:", err)
	}

	err = json.Unmarshal(data, &Artists)
	if err != nil {
		log.Fatal("Error unmarshalling JSON 1:", err)
	}

	data2, err2 := api("https://groupietrackers.herokuapp.com/api/locations")
	var location_API_Response location_API_Response_Structure

	if err2 != nil {
		log.Fatal("Error fetching API 2:", err2)
	}

	err2 = json.Unmarshal(data2, &location_API_Response)
	if err2 != nil {
		log.Fatal("Error unmarshalling JSON 2:", err2)
	}

	Locations = location_API_Response.Index

	// Récupération des relations (dates + lieux)
	data3, err3 := api("https://groupietrackers.herokuapp.com/api/relation")
	var relation_API_Response relation_API_Response_Structure

	if err3 != nil {
		log.Fatal("Error fetching API 3:", err3)
	}

	err3 = json.Unmarshal(data3, &relation_API_Response)
	if err3 != nil {
		log.Fatal("Error unmarshalling JSON 3:", err3)
	}

	Relations = relation_API_Response.Index

	// Créer la liste de tous les concerts
	AllConcerts = BuildConcertsList()
}

func api(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET error: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

// BuildConcertsList crée une liste plate de tous les concerts à partir des relations
func BuildConcertsList() []Concert {
	var concerts []Concert

	for _, relation := range Relations {
		var artistName string
		for _, artist := range Artists {
			if artist.Id == relation.Id {
				artistName = artist.Name
				break
			}
		}

		for location, dates := range relation.DatesLocations {
			for _, date := range dates {
				concerts = append(concerts, Concert{
					ArtistName: artistName,
					Location:   location,
					Date:       date,
					ArtistId:   relation.Id,
				})
			}
		}
	}

	return concerts
}

// ParseFirstAlbumDate parse la date du premier album (format: "DD-MM-YYYY")
func ParseFirstAlbumDate(dateStr string) (time.Time, error) {
	layout := "02-01-2006"
	return time.Parse(layout, dateStr)
}

// GetArtistLocations retourne toutes les locations d'un artiste
func GetArtistLocations(artistId int) []string {
	for _, relation := range Relations {
		if relation.Id == artistId {
			var locations []string
			for location := range relation.DatesLocations {
				locations = append(locations, location)
			}
			return locations
		}
	}
	return []string{}
}
