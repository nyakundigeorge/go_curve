package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Saiyan struct {
	*Person
	Power int
}

func (p *Person) introduce() {
	fmt.Printf("Hi, I am %s\n", p.Name)
}
func main_(){
	goku := &Saiyan{
		Person : &Person{"Goku"},
		Power : 9001,
	}

	goku.introduce()

}