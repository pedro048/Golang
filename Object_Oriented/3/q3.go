package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
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
		Animal: Animal{Name: "Bob", Age: 3},
	}

	var myCat = Cat{
		Animal: Animal{Name: "Garfield", Age: 4},
	}

	PetSound(myDog)
	PetSound(myCat)

}
