package main

import (
	"fmt"

	"github.com/abyo/mygoprogram/footypes"
)

/* comment */
func main() {
	fmt.Println("test")

	var fooVar footypes.Foo

	fooVar.TypeBoolean = true
	fooVar.TypeString = "hello"

	fmt.Println(fooVar)
}

// main.main()
