package main

import "fmt"

/* comment */
func main() {
	/* syntaxe */
	/*
		var list [n]type
	*/
	var list [2]int

	list[0] = 10
	list[1] = 20

	newList := [...]int{40, 50}
	bigList := [...]int{40, 50, 69, 420, 777777, 50085}

	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(list[1])

	fmt.Println(newList)
	fmt.Println(newList[0])
	fmt.Println(newList[1])

	for pos, value := range bigList {
		fmt.Printf("Position %d est égale à %d\n", pos, value)
	}
}

// main.main()
