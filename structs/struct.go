package structs

import "fmt"

type Person struct{
	name string
	age int
}

func PracticeStructs() {
	var p1 Person = Person{
		name:"Atharva",
		age:23,
	}

	fmt.Println(p1) // {Atharva 23}
	fmt.Println("Name: ",p1.name)
}