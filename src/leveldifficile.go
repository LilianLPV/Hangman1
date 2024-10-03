package Hangman

import (
	"fmt"
	"math/rand"
)

func LevelDifficile() {
	var input string
	var lettre1, lettre2 int
	var erreursdifficile int
	const maxerreursdifficile = 6

	motD, err := WordEasyRandom("motsfacile.txt")
	runeD := []rune(motD)

	longueur := len(motD)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var LettreT []string
	var MotT []string

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
			motCache[i] = runeD[i]
		} else {
			motCache[i] = '_'
		}
	}
	for {
		fmt.Println(" ")
		fmt.Println("Voici votre mot :", string(motCache))
		fmt.Println(" ")
		fmt.Println("Lettres déjà tentées :", LettreT)
		fmt.Println(" ")
		fmt.Println("Mots déjà tentés :", MotT)
		fmt.Println(" ")
		fmt.Println("Entrez une Lettre Miniscule !")
		fmt.Println(" ")
		fmt.Scanln(&input)

		input = ToLower(input)

		if !(input >= "a" && input <= "z") {
			fmt.Println(" ")
			fmt.Println("Erreur de saisie")
			fmt.Println(" ")

		}

		if len(input) > 1 {
			if input == motD {
				fmt.Println(" ")
				fmt.Println("Félicitations ! Vous avez deviné le mot :", motD)
				fmt.Println(" ")
				Start()
				return
			} else {
				erreursdifficile += 2
				fmt.Println(" ")
				fmt.Println("Désolé,", input, "n'est pas le bon mot.")
				fmt.Println(" ")
				ClearTerminal()
				fmt.Println(" ")
				fmt.Printf("Erreurs : %d/%d\n", erreursdifficile, maxerreursdifficile)
				ErreursHangmanFacile(erreursdifficile)
				if !contains(MotT, input) {
					MotT = append(MotT, input)
				}
			}
		}
		runeInput := []rune(input)[0]
		found := false
		if len(input) == 1 {
			for i := 0; i < longueur; i++ {
				if runeD[i] == runeInput {
					motCache[i] = runeInput
					found = true
				}
			}
		}
		if found {
			fmt.Println(" ")
			fmt.Println("Bravo ! La lettre", input, "est dans le mot.")
			fmt.Println(" ")
			fmt.Printf("Erreurs : %d/%d\n", erreursdifficile, maxerreursdifficile)
			if !contains(LettreT, input) {
				LettreT = append(LettreT, input)
			}
		} else if len(input) == 1 {
			erreursdifficile++
			fmt.Println(" ")
			fmt.Println("Désolé, la lettre", input, "n'est pas dans le mot.")
			fmt.Println(" ")
			fmt.Printf("Erreurs : %d/%d\n", erreursdifficile, maxerreursdifficile)
			ErreursHangmanFacile(erreursdifficile)
			if !contains(LettreT, input) {
				LettreT = append(LettreT, input)
			}
		}
		if string(motCache) == motD {
			fmt.Println(" ")
			fmt.Println("Félicitations ! Vous avez deviné le mot :", motD)
			fmt.Println(" ")
			Start()
			return
		}
		if erreursdifficile >= maxerreursdifficile {
			fmt.Println(" ")
			fmt.Println("Désolé, vous avez fait trop d'erreurs. Le mot était :", motD)
			fmt.Println(" ")
			ErreursHangmanFacile(erreursdifficile) // Affiche le pendu final si le joueur perd
			Start()
			return
		}
	}
}
