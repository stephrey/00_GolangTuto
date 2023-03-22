package main

import "fmt" //librairie with stdio

/* comment */
func main() {
	var (
		i int
		y string
		z bool
	)

	i = 3
	y = "stephane"
	z = true

	fmt.Printf("mon nombre est : %v\n", i)
	fmt.Printf("mon nombre est : %v\n", y)
	fmt.Printf("mon nombre est : %v\n", z)

	fmt.Printf("\n")

	fmt.Printf("Bonjour,\n Je suis %v\n %v %v", y, i, z)
}

// main.main()
