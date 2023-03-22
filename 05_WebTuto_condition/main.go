package main

import (
	"fmt" //librairie with stdio
	"math/rand"
	"time"
)

/* comment */
func main() {
	rand.Seed(time.Now().UnixNano())
	if x := rand.Int(); x%2 == 0 {
		fmt.Println(x, " est un nombre paire")
	} else {
		fmt.Println(x, " est un nombre impaire")
	}

	y := rand.Int() % 2
	if y == 0 {
		fmt.Println(y, " est un nombre paire")
	} else {
		fmt.Println(y, " est un nombre impaire")
	}

	age := 17

	if age > 18 {
		fmt.Println("je suis majeur")
	} else if age == 18 {
		fmt.Println("je viens tout juste d'être majeur")
	} else {
		fmt.Println("je suis mineure")
	}
}

// main.main()
