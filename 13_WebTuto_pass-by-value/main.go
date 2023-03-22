package main

import "fmt"

//pass by value
func updateA(number int) int {
	number = 5
	return number
}

func updateB(item map[string]int) {
	item["bonbon"] = 4
}

/* comment */
func main() {
	// A: string, int, bool, float, array
	number := 10
	updateA(number)
	fmt.Println(number) //10
	number = updateA(number)
	fmt.Println(number) //5

	//B
	groceryList := map[string]int{
		"prince": 6,
		"viande": 10,
	}

	updateB(groceryList)
	fmt.Println(groceryList) //is a pointer !
}

// main.main()
