package main

import (
	"fmt"
	"go-pci/chapter2/recomendations"
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.SetOutput(os.Stdout)
	peopleNames := os.Args[1:3]

	p1, err1 := recomendations.FindPersonByName(peopleNames[0])
	CheckError(err1)

	p2, err2 := recomendations.FindPersonByName(peopleNames[1])
	CheckError(err2)

	movies := recomendations.FindSimilarMovies(p1, p2)

	for _, movie := range movies {
		fmt.Println(movie.Title)
	}
}
