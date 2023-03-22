package main

import "fmt"

/* comment */
func main() {
	w := func() {
		fmt.Println("Je suis une fonction anonyme sans return.")
	}
	w()

	z := func() string {
		println("Je suis une fonction anonyme")
		return "Alex"
	}()
	println(z)

	x, y := func() (int, int) {
		println("Aucun paramètres, Deux returns")
		return 5, 7
	}()
	println(x)
	println(y)

	func(a, b int) {
		println("A * A + B * B = ", a*a+b*b)
	}(x, y)
}

// main.main()
