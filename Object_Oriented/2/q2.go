package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Age  int
}

type Cat struct {
	Name string
	Age  int
}

type Pet interface {
	Sound()
}

func (d Dog) Sound() {
	fmt.Println("Au!")
}

func (c Cat) Sound() {
	fmt.Println("Miau!")
}

func PetSound(pet Pet) {
	pet.Sound()
}

func main() {
	var myDog = Dog{
		Name: "Bob",
		Age:  3,
	}

	var myCat = Cat{
		Name: "Garfield",
		Age:  4,
	}

	PetSound(myDog)
	PetSound(myCat)

}
