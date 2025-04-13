package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func structPlay() {
	iro := person{
		"Iroshan",
		"Vidanage",
		contactInfo{
			"iroshan@home.com",
			11111,
		},
	}
	shi := person{
		firstName: "Shihan",
		lastName:  "Vidanage",
		contactInfo: contactInfo{
			email:   "shihan@home.com",
			zipCode: 11111,
		},
	}
	fmt.Println(iro, shi)

	var ss person
	ss.printp()
	ss.firstName = "Alex"
	ss.lastName = "Anderson"
	ss.contactInfo.email = "alex@home.com"
	ss.contactInfo.zipCode = 11111
	fmt.Println(ss)
	ss.printp()

	// access the memory address of this value ss - pointer
	// ssPointer := &ss
	// ssPointer.updateName("Aleck")

	// *pointer
	// access the value of the memory address - operator

	ss.updateName("Aleck")
	ss.printp()

	fmt.Println()
	changeSlice()
	changeVar()
}

func (p person) printp() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	// (*p).firstName = newFirstName
	p.firstName = newFirstName
}

func changeSlice() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	// this gets changed
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func changeVar() {
	mySlice := "Ho ho ho"
	// this doesn't get changed
	updateVar(mySlice)
	fmt.Println(mySlice)
}

func updateVar(s string) {
	s = "Bye"
}
