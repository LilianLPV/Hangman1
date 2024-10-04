package Hangman

import "fmt"

func Start() {
	
	ClearTerminal()
fmt.Println(`                                                                                                                                                 
      ___           ___           ___           ___           ___           ___           ___     
     /__/\         /  /\         /__/\         /  /\         /__/\         /  /\         /__/\    
     \  \:\       /  /::\        \  \:\       /  /:/_       |  |::\       /  /::\        \  \:\   
      \__\:\     /  /:/\:\        \  \:\     /  /:/ /\      |  |:|:\     /  /:/\:\        \  \:\  
  ___ /  /::\   /  /:/─/::\   _____\__\:\   /  /:/_/::\   __|__|:|\:\   /  /:/~/::\   _____\__\:\ 
 /__/\  /:/\:\ /__/:/ /:/\:\ /__/::::::::\ /__/:/__\/\:\ /__/::::| \:\ /__/:/ /:/\:\ /__/::::::::\
 \  \:\/:/__\/ \  \:\/:/__\/ \  \:\──\──\/ \  \:\ /──/:/ \  \:\──\__\/ \  \:\/:/__\/ \  \:\──\──\/
  \  \::/       \  \::/       \  \:\  ───   \  \:\  /:/   \  \:\        \  \::/       \  \:\  ───
   \  \:\        \  \:\        \  \:\        \  \:\/:/     \  \:\        \  \:\        \  \:\     
    \  \:\        \  \:\        \  \:\        \  \::/       \  \:\        \  \:\        \  \:\    
     \__\/         \__\/         \__\/         \__\/         \__\/         \__\/         \__\/    
	  
`)
	var input string
	fmt.Println(" ")
	fmt.Println("Bienvenue sur Hangman !")
	fmt.Println(" ")
	fmt.Println("Votre choix : ")
	fmt.Println(" ")
	fmt.Println("\033[92m1. Niveau Facile\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[93m2. Niveau Moyen\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[91m3. Niveau Difficile\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[97m0. Quitter\033[0m")
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
