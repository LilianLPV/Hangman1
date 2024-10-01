package Hangman

import (
	"fmt"
	"math/rand"
)

func LevelFacile() {
	var input string
	var lettre1, lettre2 int
	var erreursfacile int
	const maxerreursfacile = 10

	motF, err := WordEasyRandom("motsfacile.txt")
	runeF := []rune(motF)

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
	motCache := make([]rune, longueur)
	for i := 0; i < longueur; i++ {
		if i == lettre1 || i == lettre2 {
			motCache[i] = runeF[i]
		} else {
			motCache[i] = '_'
		}
	}
	fmt.Println(" ")
	fmt.Println("Voici votre mot :", string(motCache))
	fmt.Println(" ")
	fmt.Println("Entrez une Lettre Miniscule !")
	fmt.Scan(&input)

	if len(input) != 1 {
		fmt.Println("Erreur de saisie")
		return
	}
	input = ToLower(input)

	if !(input >= "a" && input <= "z") {
		fmt.Println("Erreur de saisie")
		return
	}

	runeInput := []rune(input)[0] 
    found := false
    for i := 0; i < longueur; i++ {
        if runeF[i] == runeInput {
            motCache[i] = runeInput
            found = true
        }
    }
  
	if found {
		fmt.Println("Bravo ! La lettre", input, "est dans le mot.")
	} else {
		erreursfacile++
		fmt.Println("Désolé, la lettre", input, "n'est pas dans le mot.")
		fmt.Printf("Erreurs : %d/%d\n", erreursfacile, maxerreursfacile)
		ErreursHangmanFacile(erreursfacile) 
	if string(motCache) == motF {
		fmt.Println("Félicitations ! Vous avez deviné le mot :", motF)
		return
	}
}
	if erreursfacile >= maxerreursfacile {
		fmt.Println("Désolé, vous avez fait trop d'erreurs. Le mot était :", motF)
		ErreursHangmanFacile(erreursfacile) // Afficher le pendu final si le joueur perd
}	
	}