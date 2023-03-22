package main

import "fmt" //librairie with stdio

/* comment */
func main() {
	var (
		x int
		y int
	)

	x = 7
	y = 3

	// + - / * %
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x / y)
	fmt.Println(x * y)
	fmt.Println(x % y)

	// == != < <= > >=
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	// && ||
	fmt.Println(x == y && x != y)
	fmt.Println(x == y || x != y)
}

// main.main()
