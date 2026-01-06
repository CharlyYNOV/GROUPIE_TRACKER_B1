package main

import (
	"fmt"
	"groupie_tracker/internals"
	controler "groupie_tracker/internals/Controler"
	"log"
	"net/http"
	"os"
)

func main() {
	internals.Main_api()

	// Routes HTML
	http.HandleFunc("/", controler.HomePage)
	http.HandleFunc("/acceuil.html", controler.HomePage)
	http.HandleFunc("/artists", controler.ViewAllArtistsPage)
	http.HandleFunc("/artist", controler.ArtistPage)
	http.HandleFunc("/concerts", controler.Concerts)
	http.HandleFunc("/search", internals.SearchBar)

	// Fichiers statiques
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Permet de remplacer le port
	port := os.Getenv("PORT")
	if port == "" {
		port = "5500"
	}

	fmt.Printf("Serveur démarré sur http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
