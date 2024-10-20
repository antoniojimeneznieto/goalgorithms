package main

import "fmt"

// Upper case means public (exported to external packages) applicable for struct fields also
// Lower case means private (not exported to external packages) applicable for struct fields also
type Person struct {
	Name     string
	Age      uint8
	employed bool
	employee Employee
}

func (p Person) agePlusOne() uint8 {
	return p.Age + 1
}

func (p *Person) incrementAge() {
	p.Age++
}

type Employee struct {
	Company string
}

func main() {
	var myPerson Person
	fmt.Println(myPerson)

	var myPerson2 Person = Person{"John", 30, true, Employee{"Google"}}
	fmt.Println(myPerson2)

	fmt.Println(myPerson2.agePlusOne())

	myPerson2.incrementAge()
	fmt.Println(myPerson2)

}
