package main

import "fmt"

type Person struct {
	name string
	age  int
}

//details func
// func printPersonDetails(psn Person) {
// 	fmt.Println("person name", psn.name)
// 	fmt.Println("person age", psn.age)

// }


//receiver function
func (psn Person) printPersonDetails() {
	fmt.Println("person name", psn.name)
	fmt.Println("person age", psn.age)
}

func main() {
	fmt.Println("Welcome Home Guys!!!!!!!")

	var person Person
	var person2 Person
	person.name = "Jakuan"
	person.age = 21
	person2 = Person{
		name: "Kuan",
		age:  20,
	}
	person.printPersonDetails()
	person2.printPersonDetails()

}
