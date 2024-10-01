package Hangman

import "fmt"

func Start() {
	
fmt.Println(`					 _   _  `)                                       
fmt.Println(`					| | | | __ _ _ __   __ _ _ __ ___   __ _ _ __  `) 
fmt.Println(`					| |_| |/ _  | '_ \ / _  | '_   _ \ / _  | '_ \ `) 
fmt.Println(`					|  _  | (_| | | | | (_| | | | | | | (_| | | | |`) 
fmt.Println(`					|_| |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|`) 
fmt.Println(`									   |___/                     	   `)   
                   
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
		fmt.Println(" ")
		LevelFacile()
		fmt.Println(" ")
	case "2":
		fmt.Println("Niveau Moyen")
		fmt.Println(" ")
		LevelMoyen()
	case "3":
		fmt.Println("Niveau Difficile")
		fmt.Println(" ")
		LevelDifficile()
		fmt.Println(" ")
	case "0":
		fmt.Println("Quitter")
		fmt.Println(" ")
		fmt.Println("Merci d'avoir joué au jeu a bientôt !!")
		fmt.Println(" ")
	default:
		fmt.Println("Ne vous échappez pas !")
		fmt.Println(" ")
		Start()
		fmt.Println(" ")
	}
}
