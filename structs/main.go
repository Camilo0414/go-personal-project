package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	first_name string
	last_name  string
	//contact    contactInfo
	contactInfo
}

func main() {

	// First approach
	// alex := person{"Alex", "Anderson"}

	// Second approach
	// alex := person{first_name: "Alex", last_name: "Anderson"}

	// Third approach
	// var alex person

	// alex.first_name = "Alex"
	// alex.last_name = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		first_name: "Jim",
		last_name:  "Halperd",
		contactInfo: contactInfo{
			email: "j.halperd@gmail.com",
			zip:   95120,
		},
	}

	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()

	// &variable = give me the memory address of the value this variable is pointing at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).first_name = newFirstName
	// *pointer = give me the value this memory address is pointing at
}
