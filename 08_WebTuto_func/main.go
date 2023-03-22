package main

import (
	"errors"
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Bonjour à tous, je m'appelle %v.\n", name)
}

// func calculatePerimRect(lo int, la int)
func calculatePerimRect(lo, la int) int {
	return 2 * (lo + la)
}

func sayBye(name string) (string, error) {
	if name == "" {
		return "", errors.New("\bErreur: Vous avez oublié de spécifier un nom!")
	}
	byeMessage := fmt.Sprintf("%v s'en va! Bonne soirée!", name)
	return byeMessage, nil
}

/* comment */
func main() {
	sayHello("Alex")
	fmt.Println(calculatePerimRect(10, 12))
	fmt.Println(sayBye("Stephane"))
	fmt.Println(sayBye(""))
}

// main.main()
