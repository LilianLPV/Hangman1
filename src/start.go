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
	case "2":
		fmt.Println("Niveau Moyen")
	case "3":
		fmt.Println("Niveau Difficile")
	case "0":
		fmt.Println("Quitter")
	default:
		fmt.Println("Ne vous échappez pas !")
		Start()
	}
}
