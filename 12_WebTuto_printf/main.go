package main

import "fmt"

/* comment */
//https://pkg.go.dev/fmt
func main() {
	//amener à disparaitre dans le futur
	println("Salut à tous")
	print("Salut à tous\n")
	//à utiliser (fmt)
	fmt.Println("Salut à tous")
	fmt.Print("Salut à tous\n")

	//printf
	x, y := 50, 35
	fmt.Printf("Mon nombre (defaut) est :%v\n", x)
	fmt.Printf("Mon nombre (base 2) est :%b\n", x)
	fmt.Printf("Mon nombre (base 8) est :%o\n", x)
	fmt.Printf("Mon nombre (base 16) est :%x\n", x)

	fmt.Printf("La valeur de X est égale à la valeur de Y -> %t\n", x == y)
	fmt.Printf("L'écriture scientifique de 1245.7890 -> :%e\n", 145.7890)
}

// main.main()
