package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json: "name"`
	Age int `json: "age"`
	Email string `json: "email"`
	Phone string `json: "phone"`
}

func main() {
	jsonFomApi := `
	[
		{
			"name": "Tsiaro",
			"age": 42,
			"email": "foo@gmail.com",
			"phone": "12346789"
		},
		{
			"name": "Justine",
			"age": 35,
			"email": "foo@gmail.com",
			"phone": "10111234678"
		}
	]
	`

	var users []User

	err := json.Unmarshal([]byte(jsonFomApi), &users)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
	fmt.Printf("json: %v\n", users)

	fmt.Println("--------------------------------------")

	var myStruct []User

	var user_one User
	user_one.Name = "Pierre"
	user_one.Age  = 31
	user_one.Email = "foo@gmail.com"
	user_one.Phone = "23456782"

	myStruct = append(myStruct, user_one)

	var user_two User
	user_two.Name = "Justin"
	user_two.Age  = 29
	user_two.Email = "foo@gmail.com"
	user_two.Phone = "563456782"

	myStruct = append(myStruct, user_two)

	jsonFromSlice, err := json.MarshalIndent(myStruct, "", " ")
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
	fmt.Println(string(jsonFromSlice))
}
