package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

var Artists []Artist
var Locations []Location

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
}

func api(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("HTTP GET error: %w", err)
	}

	defer response.Body.Close() // sert a libérer les ressources allouées à la première ligne du programme et a fermer la connexion active

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body) // check dans le body si il y a une erreur (ne s'arrête pas de lire tant qu'il n'y a pas d'erreur ou EOF)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}
