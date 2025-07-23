package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	age int
	contactInfo
}

func main() {
	shahathir := person{
		firstName: "Shahathir", 
		lastName: "Iskandar", 
		age: 23, 
		contactInfo: contactInfo{
			email: "shahathiriskandar42@gmail.com", 
			zipCode: 68100,
		},
	}
	fmt.Println("Hello, my name is", shahathir.firstName, shahathir.lastName,"and I am", shahathir.age, "years old")
	fmt.Println("My email is", shahathir.contactInfo.email, "and my zip code is", shahathir.contactInfo.zipCode)

	fmt.Println("Before update:")
	shahathir.print()
	fmt.Println()

	fmt.Println("After update:")
	shahathirPointer := &shahathir
	shahathirPointer.updateName("Shah")
	shahathir.print()
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}