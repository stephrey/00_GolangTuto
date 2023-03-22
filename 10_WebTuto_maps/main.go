package main

import "fmt"

/* comment */
func main() {
	supermarketPrice := map[string]int{
		"prince": 8,
		"eau":    2,
		"meat":   6,
	}
	fmt.Println(supermarketPrice["prince"])

	for key, value := range supermarketPrice {
		fmt.Println(key, value)
	}

	DaysInAYear := 0
	DaysInaMounth := map[string]int{
		"janvier":   31,
		"février":   28,
		"mars":      31,
		"avril":     30,
		"mai":       31,
		"juin":      30,
		"juillet":   31,
		"aout":      31,
		"septembre": 30,
		"octobre":   31,
		"novembre":  30,
		"décembre":  31,
	}
	for key, value := range DaysInaMounth {
		fmt.Println(key, value)
		DaysInAYear = DaysInAYear + value
	}
	fmt.Println(DaysInAYear)
}

// main.main()
