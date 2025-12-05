package main

import (
	"fmt"
	"groupie_tracker/internals"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Page d'accueil
func homePage(w http.ResponseWriter, r *http.Request) {
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

	data := struct {
		Artists   []internals.Artist
		Locations []internals.Location
	}{
		Artists:   displayArtists,
		Locations: displayLocations,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}

// Nouvelle route pour la page des Artistes
func viewAllArtistsPage(w http.ResponseWriter, r *http.Request) {
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

	if search_input != "" {
		internals.Artists = internals.FilterArtists(internals.Artists, search_input)
	}

	data := struct{ Artists []internals.Artist }{Artists: internals.Artists}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}

func concerts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		w.WriteHeader(http.StatusOK)
		return
	}

	tmpl, err := template.ParseFiles("./templates/concerts.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Error while loading page", http.StatusInternalServerError)
		return
	}

	data := struct{ Artists []internals.Artist }{Artists: internals.Artists}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}

func main() {
	internals.Main_api()

	// Routes HTML
	http.HandleFunc("/", homePage)
	http.HandleFunc("/acceuil.html", homePage)
	http.HandleFunc("/artists", viewAllArtistsPage)
	http.HandleFunc("/concerts", concerts)
	http.HandleFunc("/search", internals.SearchBar)

	// Fichiers statiques
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))

	// Permet de remplacer le port
	port := os.Getenv("PORT")
	if port == "" {
		port = "5500"
	}

	fmt.Printf("Serveur démarré sur http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
