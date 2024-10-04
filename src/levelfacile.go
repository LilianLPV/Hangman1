package Hangman

import (
	"fmt"
	"math/rand"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func LevelFacile() {
	var input string
	var lettre1, lettre2 int
	var erreursfacile int
	const maxerreursfacile = 10

	motF, err := WordEasyRandom("docs/motsfacile.txt")
	runeF := []rune(motF)

	longueur := len(motF)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var LettreT []string
	var MotT []string
	ClearTerminal()
	fmt.Println("Bienvenue dans le \033[92mNiveau Facile\033[0m !")
	fmt.Println(" ")
	fmt.Println("\033[1m\033[45mAttention les lettres afficher peuvent être afficher plusieurs fois !\033[0m")
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
	for {
		fmt.Println(" ")
		fmt.Println("Voici votre mot :", string(motCache))
		fmt.Println(" ")
		fmt.Println("Lettres déjà tentées :", LettreT,)
		fmt.Println(" ")
		fmt.Println("Mots déjà tentés :", MotT,)
		fmt.Println(" ")
		fmt.Println("Entrez une Lettre Miniscule ou un Mot !")
		fmt.Println(" ")
		fmt.Scanln(&input)

		input = ToLower(input)

		if !(input >= "a" && input <= "z") {
			fmt.Println(" ")
			fmt.Println("Erreur de saisie, veuillez mettre une lettre ou un mot !")
			fmt.Println(" ")
		}

		if len(input) > 1 {
			if input == motF {
				ClearTerminal()
				fmt.Println(" ")
				fmt.Println("Félicitations ! Vous avez deviné le mot :", motF)
				fmt.Println(" ")
				Start()
				return
			} else {
				erreursfacile += 2
				ClearTerminal()
				fmt.Println(" ")
				fmt.Println("Désolé,", input, "n'est pas le bon mot.")
				fmt.Println(" ")
				fmt.Println(" ")
				fmt.Printf("Erreurs : %d/%d\n", erreursfacile, maxerreursfacile)
				ErreursHangmanFacile(erreursfacile)
				if !contains(MotT, input) {
					MotT = append(MotT, input)
				}
			}
		}
		runeInput := []rune(input)[0]
		found := false
		if len(input) == 1 {
			for i := 0; i < longueur; i++ {
				if runeF[i] == runeInput {
					motCache[i] = runeInput
					found = true
				}
			}
		}
		if found {
			ClearTerminal()
			fmt.Println(" ")
			fmt.Println("Bravo ! La lettre", input, "est dans le mot.")
			fmt.Println(" ")
			fmt.Printf("Erreurs : %d/%d\n", erreursfacile, maxerreursfacile)
			if !contains(LettreT, input) {
			LettreT = append(LettreT, input)
			}
		} else if len(input) == 1 {
			erreursfacile++
			ClearTerminal()
			fmt.Println(" ")
			fmt.Println("Désolé, la lettre", input, "n'est pas dans le mot.")
			fmt.Println(" ")

			fmt.Printf("Erreurs : %d/%d\n", erreursfacile, maxerreursfacile)
			ErreursHangmanFacile(erreursfacile)
			if !contains(LettreT, input) {
			LettreT = append(LettreT, input)
			}
		}
		if string(motCache) == motF {
			ClearTerminal()
			fmt.Println(" ")
			fmt.Println("Félicitations ! Vous avez deviné le mot :", motF)
			fmt.Println(" ")
			Start()
			return
		}
		if erreursfacile >= maxerreursfacile {
			ClearTerminal()
			fmt.Println(" ")
			fmt.Println("Désolé, vous avez fait trop d'erreurs. Le mot était :", motF)
			fmt.Println(" ")
			ErreursHangmanFacile(erreursfacile) // Affiche le pendu final si le joueur perd
			Start()
			return
		}
	}
}
