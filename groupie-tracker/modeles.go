package main

// ----- Données brutes renvoyées par l'API -----

type Artiste struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Nom          string   `json:"name"`
	Membres      []string `json:"members"`
	DateCreation int      `json:"creationDate"`
	PremierAlbum string   `json:"firstAlbum"`
}

type Lieux struct {
	ID    int      `json:"id"`
	Lieux []string `json:"locations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID         int                 `json:"id"`
	DatesLieux map[string][]string `json:"datesLocations"`
}

// ----- Enveloppes : locations, dates et relation sont dans une clé "index" -----

type ReponseLieux struct {
	Index []Lieux `json:"index"`
}

type ReponseDates struct {
	Index []Dates `json:"index"`
}

type ReponseRelation struct {
	Index []Relation `json:"index"`
}

// ----- Structure complète envoyée aux templates HTML -----

type ArtisteComplet struct {
	Artiste
	Lieux      []string
	Dates      []string
	DatesLieux map[string][]string
}
