package main

import (
	"fmt"
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

	http.ServeFile(w, r, "./templates/accueil.html")
}

func main() {
	// Routes HTML
	// la racine (/) sert la page d'accueil ; garder le chemin explicite pour compatibilité
	http.HandleFunc("/", homePage)
	http.HandleFunc("/acceuil.html", homePage)

	// Fichiers statiques
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("./static/script"))))
	// Images et autres ressources statiques
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))

	// Permet de remplacer le port via la variable d'environnement PORT (utile pour l'hébergement ou éviter les conflits)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5500"
	}

	fmt.Printf("Serveur démarré sur http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
