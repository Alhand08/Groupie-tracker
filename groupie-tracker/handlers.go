package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
)

// donneesErreur est envoyée au template erreur.html.
type donneesErreur struct {
	Code    int
	Message string
}

// rendre génère la page dans un tampon avant de l'envoyer :
// si le template échoue, on n'envoie pas une page à moitié écrite.
func rendre(w http.ResponseWriter, nomTemplate string, donnees any) {
	var tampon bytes.Buffer
	if err := modeles.ExecuteTemplate(&tampon, nomTemplate, donnees); err != nil {
		log.Printf("erreur template %s : %v", nomTemplate, err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tampon.WriteTo(w)
}

// pageErreur affiche une page d'erreur propre avec le bon code HTTP.
func pageErreur(w http.ResponseWriter, code int, message string) {
	var tampon bytes.Buffer
	err := modeles.ExecuteTemplate(&tampon, "erreur.html", donneesErreur{Code: code, Message: message})
	if err != nil {
		log.Printf("erreur template erreur.html : %v", err)
		http.Error(w, message, code)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	tampon.WriteTo(w)
}

// pageAccueil : liste de tous les artistes (URL : /)
func pageAccueil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pageErreur(w, http.StatusNotFound, "Page introuvable")
		return
	}
	if r.Method != http.MethodGet {
		pageErreur(w, http.StatusMethodNotAllowed, "Méthode non autorisée")
		return
	}
	rendre(w, "index.html", listeArtistes)
}

// pageArtiste : détail d'un artiste (URL : /artiste?id=3)
func pageArtiste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pageErreur(w, http.StatusMethodNotAllowed, "Méthode non autorisée")
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		pageErreur(w, http.StatusBadRequest, "Identifiant d'artiste invalide")
		return
	}

	for _, artiste := range listeArtistes {
		if artiste.ID == id {
			rendre(w, "artiste.html", artiste)
			return
		}
	}
	pageErreur(w, http.StatusNotFound, "Artiste introuvable")
}
