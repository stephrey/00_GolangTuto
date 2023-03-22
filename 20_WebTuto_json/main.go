package main

//https://www.youtube.com/watch?v=DReUmZ_VyPI&list=PLuWyq_EO5_AKP_KCaIr53UfOqlPThTXat&index=21

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

/* comment */
func main() {
	jsonFromApi := `
		[
			{
				"name": "Alex",
				"age":	23,
				"email": "foo@gmail.com",
				"phone": "123456789"
			},
			{
				"name": "Justine",
				"age":	32,
				"email": "bar@gmail.com",
				"phone": "100011101234"
			}
		]`

	//read from json
	//fmt.Println(jsonFromApi)
	var users []User //we use a slide (dynamic array) why the json length is not know
	err := json.Unmarshal([]byte(jsonFromApi), &users)
	if err != nil {
		fmt.Println("Error unmarshalling json: ", err)
	} else {
		fmt.Println(users)
	}

	//write in json and read back
	var myStruct []User
	var user_one User
	user_one.Name = "Pierre"
	user_one.Age = 31
	user_one.Email = "pierre@fleur.com"
	user_one.Phone = "+41266632020"
	myStruct = append(myStruct, user_one)

	var user_two User
	user_two.Name = "Daniel"
	user_two.Age = 51
	user_two.Email = "daniel@fleur.com"
	user_two.Phone = "+41266632022"
	myStruct = append(myStruct, user_two)

	//JsonFromSlice, err := json.Marshal(myStruct)
	JsonFromSlice, err := json.MarshalIndent(myStruct, "", " ")
	if err != nil {
		fmt.Println("Error marshalling : ", err)
	} else {
		fmt.Println(string(JsonFromSlice))
	}
}

// main.main()
