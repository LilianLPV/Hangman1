package Hangman

import "fmt"

func ErreursHangmanFacile(erreurs int) {
    // Liste des étapes du pendu, sous forme de chaîne de caractères
    pendu := []string{
        `
    _______
    |    
    |    
    |    
    |    
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |    
    |    
    |    
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |     O
    |    
    |    
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |     O
    |     |
    |    
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |     O
    |    /|
    |    
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |     O
    |    /|\
    |    
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |     O
    |    /|\
    |    / 
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |     O
    |    /|\
    |    / \
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |     O
    |   --|--
    |    / \
    |    
  __|__
  |    |______
  |___________|
        `,
        `
    _______
    |     |
    |    [O]
    |   --|--
    |    / \
    |    
  __|__
  |    |______
  |___________|
        `,
    }

    // Afficher l'étape correspondant au nombre d'erreurs
    if erreurs < len(pendu) {
        fmt.Println(pendu[erreurs])
    } else {
        fmt.Println("Nombre d'erreurs dépassé.")
    }
}