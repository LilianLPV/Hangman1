package Hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func WordHardRandom(motsdifficile string) (string, error) {
	file, err := os.Open(motsdifficile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var motD []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		motD = append(motD, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(motD) == 0 {
		return "", fmt.Errorf("le fichier est vide ou n'a pas de mots valides")
	}

	rand.Seed(time.Now().UnixNano())
	motAleatoire := motD[rand.Intn(len(motD))]

	return motAleatoire, nil
}


