package main

import (
	"html/template"
	"log"
	"net/http"
)

// Données et templates chargés une seule fois au démarrage.
var (
	listeArtistes []ArtisteComplet
	modeles       *template.Template
)

func main() {
	var err error

	// 1. Charger les templates HTML
	modeles, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("impossible de charger les templates : %v", err)
	}

	// 2. Récupérer les données de l'API (une seule fois)
	log.Println("Chargement des données de l'API...")
	listeArtistes, err = chargerDonnees()
	if err != nil {
		log.Fatalf("impossible de charger les données : %v", err)
	}
	log.Printf("%d artistes chargés", len(listeArtistes))

	// 3. Définir les routes
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/style.css")
	})
	http.HandleFunc("/", pageAccueil)
	http.HandleFunc("/artiste", pageArtiste)

	// 4. Lancer le serveur
	log.Println("Serveur lancé sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("le serveur s'est arrêté : %v", err)
	}
}
