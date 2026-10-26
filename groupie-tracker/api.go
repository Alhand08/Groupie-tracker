package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Vérifie cette URL avec celle fournie par ton prof.
const urlAPI = "https://groupietrackers.herokuapp.com/api"

var clientHTTP = &http.Client{Timeout: 10 * time.Second}

// recupererJSON appelle une URL et remplit "cible" avec le JSON reçu.
func recupererJSON(url string, cible any) error {
	reponse, err := clientHTTP.Get(url)
	if err != nil {
		return fmt.Errorf("appel de %s : %w", url, err)
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != http.StatusOK {
		return fmt.Errorf("appel de %s : statut %d", url, reponse.StatusCode)
	}

	if err := json.NewDecoder(reponse.Body).Decode(cible); err != nil {
		return fmt.Errorf("décodage de %s : %w", url, err)
	}
	return nil
}

// chargerDonnees récupère les 4 parties de l'API et les fusionne par ID.
func chargerDonnees() ([]ArtisteComplet, error) {
	var artistes []Artiste
	var reponseLieux ReponseLieux
	var reponseDates ReponseDates
	var reponseRelation ReponseRelation

	if err := recupererJSON(urlAPI+"/artists", &artistes); err != nil {
		return nil, err
	}
	if err := recupererJSON(urlAPI+"/locations", &reponseLieux); err != nil {
		return nil, err
	}
	if err := recupererJSON(urlAPI+"/dates", &reponseDates); err != nil {
		return nil, err
	}
	if err := recupererJSON(urlAPI+"/relation", &reponseRelation); err != nil {
		return nil, err
	}

	// On indexe chaque partie par ID pour les retrouver facilement.
	lieuxParID := make(map[int][]string)
	for _, l := range reponseLieux.Index {
		lieuxParID[l.ID] = l.Lieux
	}

	datesParID := make(map[int][]string)
	for _, d := range reponseDates.Index {
		datesParID[d.ID] = d.Dates
	}

	relationParID := make(map[int]map[string][]string)
	for _, r := range reponseRelation.Index {
		relationParID[r.ID] = r.DatesLieux
	}

	resultat := make([]ArtisteComplet, 0, len(artistes))
	for _, a := range artistes {
		resultat = append(resultat, ArtisteComplet{
			Artiste:    a,
			Lieux:      lieuxParID[a.ID],
			Dates:      datesParID[a.ID],
			DatesLieux: relationParID[a.ID],
		})
	}
	return resultat, nil
}
