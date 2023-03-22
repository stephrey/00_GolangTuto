package main

// go run .\main.go .\contact.go

type contact struct {
	name  string
	age   int
	phone map[string]string
}

func newContact(name string, age int) contact {
	c := contact{
		name:  name,
		age:   age,
		phone: map[string]string{},
	}
	return c
}
