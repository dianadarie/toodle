package handlers

import (
	"errors"
	"fmt"
	"toodle/game"
)

type Handler struct {
	Toodle game.Toodle
}

func NewHandler() (Handler, error) {
	toodle, err := game.NewToodle()
	if err != nil {
		fmt.Printf("%v", err)
	}

	return Handler{Toodle: toodle}, nil
}

func (h *Handler) AddUserGuess(username, password, word string) (string, error) {
	// Check user exists
	ok, err := h.Toodle.SheetsClient.IsUser(username, password)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("user doesn't exist, password is wrong or user is not approved")
	}

	// Check user no of attempts
	
	// Check input word against word of the day
	guess, err := h.Toodle.Play(word)
	if err != nil {
		return "", err
	}

	// Increase no of attempts

	// Check if user is winner

	// Check game over

	fmt.Printf("%s", guess)
	return "", nil
}
