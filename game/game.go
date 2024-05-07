package game

import (
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"toodle/game/sheets"
)

type Toodle struct {
	Dictionary   map[string]bool
	GameWords    map[string]bool
	WordOfDay    string
	SheetsClient sheets.GoogleSheets
}

func NewToodle() (Toodle, error) {
	dictionary, err := loadWords("game/words/accepted_words.txt")
	if err != nil {
		return Toodle{}, err
	}

	gameWords, err := loadWords("game/words/game_words.txt")
	if err != nil {
		return Toodle{}, err
	}

	sClient, err := sheets.NewSheetsClient()
	if err != nil {
		return Toodle{}, err
	}
	return Toodle{dictionary, gameWords, "", sClient}, nil
}

func (t *Toodle) GetWordOfTheDay(sheetsClient sheets.GoogleSheets) (string, error) {
	return "", nil
}

func (t *Toodle) Play(attempt string) (string, error) {
	word, err := t.SheetsClient.GetWordGivenDay()
	if err != nil {
		return "", err
	}

	if word == "" {
		word = t.getRandomWord()
		err = t.SheetsClient.SetWordGivenDay(word)
		if err != nil {
			return "", err
		}
	}

	attempt = strings.ToUpper(attempt)
	_, ok := t.Dictionary[attempt]
	if !ok {
		return "", errors.New("not a proper word")
	}

	guess := ""
	for i := 0; i < len(attempt); i++ {
		if attempt[i] == word[i] {
			guess += "2"
		} else if containsChar(word, rune(attempt[i])) {
			guess += "1"
		} else {
			guess += "0"
		}
	}
	return guess, nil
}

func (t *Toodle) getRandomWord() string {
	randomIndex := rand.Intn(len(t.GameWords))

	// Iterate over the map to find the word at the random index
	index := 0
	for word := range t.GameWords {
		if index == randomIndex {
			return word
		}
		index++
	}

	return ""
}

func loadWords(path string) (map[string]bool, error) {
	pwd, _ := os.Getwd()
	file, err := os.Open(filepath.Join(pwd, path))
	if err != nil {
		return nil, errors.New("error reading dictionary file")
	}
	defer file.Close()

	wordMap := make(map[string]bool)
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			break // end of file
		}

		// Split the line into words
		words := strings.Fields(string(buffer[:n]))

		// Add each word to the map
		for _, word := range words {
			wordMap[word] = true
		}
	}

	return wordMap, nil
}

func containsChar(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}
