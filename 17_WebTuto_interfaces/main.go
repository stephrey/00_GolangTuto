package main

import "fmt"

type intfAnimal interface {
	Noise() string
	NumberOfLegs() int
}

type Cat struct {
	Name  string
	Breed string
}

type Spider struct {
	Name     string
	Breed    string
	Venomous bool
}

/* comment */
func main() {
	cat := Cat{
		Name:  "Kitty",
		Breed: "Siamois",
	}
	PrintAnimalInfo(&cat)

	spyder := Spider{
		Name:     "Spyddy",
		Breed:    "none",
		Venomous: true,
	}
	PrintAnimalInfo(&spyder)
}

func PrintAnimalInfo(a intfAnimal) {
	fmt.Println("Le cris de cet animal est ", a.Noise(), "et il possède ", a.NumberOfLegs(), "pattes")
}

func (c *Cat) Noise() string {
	return "miaou"
}

func (c *Cat) NumberOfLegs() int {
	return 4
}

func (c *Spider) Noise() string {
	return "hiiis"
}

func (c *Spider) NumberOfLegs() int {
	return 8
}

// main.main()
