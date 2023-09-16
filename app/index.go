package app

import (
	"fmt"
)


type dog struct {
	name, breed string
}

func (d dog) bark() {
	fmt.Printf("%v is barking", d.name)
}

func Bark() {
	newDog := dog{"daisy", "german-shepard"}
	newDog.bark()

	newD := dog{"goerge", "pitbull"}
	newD.bark()

	new := dog{"marcuz", "german-shepard"}
	new.bark()
}