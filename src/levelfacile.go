package Hangman

import (
	"fmt"
	"math/rand"
)

func LevelFacile() {
	var lettre1, lettre2 int
	motF, err := WordEasyRandom("motsfacile.txt")
	longueur := len(motF)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("Bienvenue dans le Niveau Facile !")
	lettre1 = rand.Intn(longueur)
	for {
		lettre2 = rand.Intn(longueur)
		if lettre1 != lettre2 {
			break
		}
	}
	motCache := ""
    for i := 0; i < longueur; i++ {
        if i == lettre1 || i == lettre2 {
            motCache += string(motF[i])
        } else {
            motCache += "_"
        }
    }
	fmt.Println(" ")
	fmt.Println("Voici votre mot :", motCache)
	fmt.Println(" ")
}

