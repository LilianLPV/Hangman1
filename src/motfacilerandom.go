package Hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func WordEasyRandom(motsfacile string) (string, error) {
	file, err := os.Open(motsfacile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var motF []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		motF = append(motF, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(motF) == 0 {
		return "", fmt.Errorf("le fichier est vide ou n'a pas de mots valides")
	}

	rand.Seed(time.Now().UnixNano())
	motAleatoire := motF[rand.Intn(len(motF))]

	return motAleatoire, nil
}


