package Hangman

import "fmt"

func ErreursHangmanFacile(erreurs int) {
    // Liste des étapes du pendu
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
    |     ☺
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
    |     ☺
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
    |     ☺
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
    |     ☺
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
    |     ☺
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
    |     ☺
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
    |     ☺
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
    |    [☺]
    |   --|--
    |    / \
    |    
  __|__
  |    |______
  |___________|
        `,
    }

    // Affichage l'étape correspondant au nombre d'erreurs
    if erreurs < len(pendu) {
        fmt.Println(pendu[erreurs])
    } else {
        fmt.Println("Nombre d'erreurs dépassé.")
    }
}

func ErreursHangmanMoyen(erreurs int) {
  // Liste des étapes du pendu
  pendu := []string{

      `
  _______
  |     |
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |    [☺]
  |   --|--
  |    / \
  |    
__|__
|    |______
|___________|
      `,
  }

  // Affichage l'étape correspondant au nombre d'erreurs
  if erreurs < len(pendu) {
      fmt.Println(pendu[erreurs])
  } else {
      fmt.Println("Nombre d'erreurs dépassé.")
  }
}

func ErreursHangmanDifficile(erreurs int) {
  // Liste des étapes du pendu
  pendu := []string{
    
      `
  _______
  |     |
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |     ☺
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
  |    [☺]
  |   --|--
  |    / \
  |    
__|__
|    |______
|___________|
      `,
  }

  // Affichage l'étape correspondant au nombre d'erreurs
  if erreurs < len(pendu) {
      fmt.Println(pendu[erreurs])
  } else {
      fmt.Println("Nombre d'erreurs dépassé.")
  }
}