package main

import (
	"fmt"
	"github.com/atilasantos/go-banksimulation/apps/producer/src/model/person"
)

func main() {
	p1 := new(person.Person)
	p1.SetAge(34)
	p1.SetName("John")

	fmt.Println(p1)

}
