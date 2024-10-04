package Hangman

import "fmt"

func Start() {
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
