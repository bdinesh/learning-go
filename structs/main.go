package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	dinesh := person{
		firstName: "Dinesh",
		lastName:  "Bogolu",
		contactInfo: contactInfo{
			email:   "abc@example.com",
			zipCode: 12345,
		},
	}

	// fmt.Println((&dinesh))
	dinesh.updateFirstName("Dinu")
	dinesh.print()
}

// A receiver function using pointers. Go always pass arguments by value.
// In order to update a struct we need to use pointers
// else the update won't be reflected in the value that we passed to the function.
func (p *person) updateFirstName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
