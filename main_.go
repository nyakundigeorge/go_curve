package main

import (
	"fmt"
)

type Saiyan struct{
	Name string
	Power int
}

func main_(){
	goku := &Saiyan{"Goku",9000}
	Super(goku)
	fmt.Println(goku.Power)
	goku.getAway()
}

func Super(s *Saiyan){
	s.Power +=10000
}

func (s *Saiyan) getAway(){
	fmt.Println("hehehehhehehe")
}

func NewSaiyan(name string, power int) *Saiyan{
	return &Saiyan{
		Name : name,
		Power : power,
	}
}