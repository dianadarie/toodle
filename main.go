package main

import (
	"fmt"
	"log"
	"toodle/handlers"
)

func handleRequests() {
	h, err := handlers.NewHandler()
	if err != nil {
		log.Fatalf("error creating handler %v", err)
	}

	guess, err := h.AddUserGuess("dianadarie", "somepass", "dream")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", guess)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
