package main

import (
	"fmt"

	"github.com/abyo/mygoprogram/functions"
)

func CalculateValue(intChan chan int) {
	randomNumer := functions.GenerateRandomNumber(59)
	intChan <- randomNumer
}

/* comment */
func main() {
	foo := make(chan int)
	defer close(foo)

	go CalculateValue(foo)
	num := <-foo
	fmt.Println(num)
	fmt.Println(foo) // get the cahnnel pointer back
}

// main.main()
