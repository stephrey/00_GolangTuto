package main

import "fmt"

var x int //global var

/* comment */
func main() {
	x = 5 //global var
	fmt.Println(x)
	f()
	fmt.Println(x)
	showY()
}

func f() {
	x := 10 //local var
	fmt.Println(x)
}

// main.main()
