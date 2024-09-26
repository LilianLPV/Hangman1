package Hangman

import (
	"fmt"
	"math/rand"
)

func LevelMoyen() {
	var pos1, pos2 int
	motM, err := WordMediumRandom("motsmoyen.txt")
	longueur := len(motM)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("Bienvenue dans le Niveau Moyen !")
	pos1 = rand.Intn(longueur)
	for {
		pos2 = rand.Intn(longueur)
		if pos1 != pos2 {
			break
		}
	}
	motCache := ""
	for i := 0; i < longueur; i++ {
		if i == pos1 || i == pos2 {
			motCache += string(motM[i])
		} else {
			motCache += "_"
		}
	}
	fmt.Println(" ")
	fmt.Println("Voici votre mot :", motCache)
	fmt.Println(" ")
}
