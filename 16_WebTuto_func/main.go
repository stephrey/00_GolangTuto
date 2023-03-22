package main

import "fmt"

//go run .\main.go .\contact.go

/* comment */
func main() {
	// func newContact return a type contact
	myContact := newContact("alex", 30, map[string]string{"fixe": "6632019", "portable": "0796758720"})
	fmt.Println(myContact)
	fmt.Println(myContact.displayContact())

	myNewContact := newContact("justine", 28, map[string]string{"fixe": "6633770", "portable": "", "fax": "0266633777"})
	fmt.Println(myNewContact)
	fmt.Println(myNewContact.displayContact())
}

// main.main()
