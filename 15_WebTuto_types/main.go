package main

//go run .\main.go .\contact.go

import "fmt"

type Example struct {
	a string
	b int
	c bool
}

/* comment */
func main() {
	myExample := Example{
		a: "alex",
		c: true,
	}
	fooExample := Example{"thomas", 12, true}
	fmt.Println(myExample)
	fmt.Println(fooExample)

	myContact := newContact("Justine", 28)
	fmt.Println(myContact)

	myContact.phone["natel"] = "+41796758720"
	myContact.phone["fax"] = "+41266632019"
	fmt.Println(myContact)
	/*
		for strucVal, val := range myExample {
			fmt.printf("ma structure contient la variable %v et sa valeur est %v\n", strucVal, val)
		}
	*/
}

// main.main()
