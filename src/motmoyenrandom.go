package Hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func WordMediumRandom(motsmoyen string) (string, error) {
	file, err := os.Open(motsmoyen)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var motM []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		motM = append(motM, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(motM) == 0 {
		return "", fmt.Errorf("le fichier est vide ou n'a pas de mots valides")
	}
	rand.Seed(time.Now().UnixNano())
	motAleatoire := motM[rand.Intn(len(motM))]
	return motAleatoire, nil
}
