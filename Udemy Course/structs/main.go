package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// placing a struct as required into a seperate struct
	contact contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {

	// // declaring struct
	// alex := person{firstName: "alex", lastName: "field"}
	// //amending data in struct
	// alex.firstName = "NEW ALEX"
	// print(alex.firstName, " ", alex.lastName)

	jim := person{
		firstName: "Jim",
		lastName:  "Blogs",
		contact: contactInfo{
			email: "jim@gmail.com",
			zip:   02342,
		},
	}
	//jim.print()

	//this is a go lang pointer it points to the address where the jim object is stored in memory
	jimPointer := &jim

	//we now use the pointer instead of the object to perform CRUD operations on
	jimPointer.updateName("Joe")

	jim.print()
}

// func dosnt work due to lack of pointers go is a pass by value language
func (pointerToPerson *person) updateName(newFName string) {
	(*pointerToPerson).firstName = newFName
}
func (p person) print() {
	fmt.Printf("%+v", p)

}

// notes  turn Address into value with *address
// turn Value into address with &value
