package main

import "fmt"

//pass by pointer
func update(PointerOfNumber *int, value int) {
	*PointerOfNumber = value
}

/* comment */
func main() {
	// A: string, int, bool, float, array
	number := 10
	fmt.Println(number) //10
	fmt.Println("Adresse mémoire de number: ", &number)

	myPointer := &number
	fmt.Println(myPointer)

	fmt.Printf("la valeur de l'addressse mémoire %v est %d\n", myPointer, *myPointer)

	update(myPointer, 20)
	fmt.Println(number) //5
	fmt.Printf("la valeur de l'addressse mémoire %v est %d\n", myPointer, *myPointer)

	// pointer and maps
	ArticleList := map[string]int{
		"controler": 800,
		"terminal":  400,
		"complex":   600,
	}
	for art, val := range ArticleList {
		fmt.Printf("Adresse mémoire de %v: %v sa valeur est : %v\n", &art, art, val)
	}
}

// main.main()
