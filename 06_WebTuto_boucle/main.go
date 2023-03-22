package main

import "fmt"

/* comment */
func main() {

	/* syntax d'une boucle for */
	/*
		for InitSimpleStatement; Condition; Updatestatement   {
			//code
		}
	*/

	fmt.Println("--------------------boucle for")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------------------boucle for simple")
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	fmt.Println("--------------------boucle for avec break")
	a := 0
	for {
		if a > 5 {
			break
		}
		fmt.Println(a)
		a++
	}

	fmt.Println("--------------------boucle for avec continue")
	a = 0
	for ; a <= 10; a++ {
		if a%2 == 1 {
			continue // reprend au début de la boucle si le nombre est impaire
		}
		fmt.Println(a)
	}
}

// main.main()
