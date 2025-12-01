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
		http.Error(w, "Error while loading page", http.StatusInternalServerError)
		return
	}

	displayArtists := internals.Artists
	if len(displayArtists) > 10 {
		displayArtists = displayArtists[:10]
	}

	data := struct{ Artists []internals.Artist }{Artists: displayArtists}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}

	displayLocations := internals.Locations
	if len(displayArtists) > 3 {
		displayLocations = displayLocations[:3]
	}

	data2 := struct{ Locations []internals.Location }{Locations: displayLocations}

	err = tmpl.Execute(w, data2)
	if err != nil {
		http.Error(w, "Error while displaying page", http.StatusInternalServerError)
		return
	}
}

func main() {
	internals.Main_api()
	// Routes HTML
	// la racine (/) sert la page d'accueil ; garder le chemin explicite pour compatibilité
	http.HandleFunc("/", homePage)
	http.HandleFunc("/acceuil.html", homePage)
	http.HandleFunc("/templates/viewAllArtists.html", homePage)

	// Fichiers statiques
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("./static/script"))))
	// Images et autres ressources statiques
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))

	// Permet de remplacer le port
	port := os.Getenv("PORT")
	if port == "" {
		port = "5500"
	}

	fmt.Printf("Serveur démarré sur http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
