package recomendations

import (
	"encoding/json"
	"go-pci/chapter2/critic"
	"os"
)

const dataFile = "data/data.json"

func RetrievePeople() ([]*critic.Person, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var people []*critic.Person
	err = json.NewDecoder(file).Decode(&people)

	return people, err
}

func FindPersonByName(name string) (*critic.Person, error) {
	people, err := RetrievePeople()
	var person *critic.Person
	for _, p := range people {
		if p.Name == name {
			person = p
			break
		}
	}
	return person, err
}

func FindSimilarMovies(p1 *critic.Person, p2 *critic.Person) []*critic.Movie {
	var similarFilms []*critic.Movie

	for _, p1Movie := range p1.Movies {

		for _, p2Movie := range p2.Movies {
			if p1Movie.Title == p2Movie.Title {
				similarFilms = append(similarFilms, p1Movie)
			}
		}
	}
	return similarFilms
}
