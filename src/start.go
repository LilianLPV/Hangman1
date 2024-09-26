package Hangman

import "fmt"

func Start() {
	var input string
	fmt.Println(" ")
	fmt.Println("Bienvenue sur Hangman !")
	fmt.Println(" ")
	fmt.Println("Votre choix : ")
	fmt.Println(" ")
	fmt.Println("1. Niveau Facile")
	fmt.Println(" ")
	fmt.Println("2. Niveau Moyen")
	fmt.Println(" ")
	fmt.Println("3. Niveau Difficile")
	fmt.Println(" ")
	fmt.Println("0. Quitter")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("Niveau Facile")
		LevelFacile()
	case "2":
		fmt.Println("Niveau Moyen")
		LevelMoyen()
	case "3":
		fmt.Println("Niveau Difficile")
		LevelDifficile()
	case "0":
		fmt.Println("Quitter")
		fmt.Println(" ")
		fmt.Println("Merci d'avoir joué au jeu a bientôt !!")
	default:
		fmt.Println("Ne vous échappez pas !")
		Start()
	}
}
