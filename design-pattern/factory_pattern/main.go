package main

import (
	"fmt"

	"example-code/design-pattern/factory_pattern/concrete"
	"example-code/design-pattern/factory_pattern/constants"
	"example-code/design-pattern/factory_pattern/factory"
)

func main() {
	// init
	cat := concrete.NewCat()
	dog := concrete.NewDog()
	rat := concrete.NewRat()
	animalFactory := &factory.AnimalFactory{
		Cat: cat,
		Dog: dog,
		Rat: rat,
	}

	// factory
	processor, err := animalFactory.GetAnimalFactory(constants.CAT)
	if err != nil {
		fmt.Println(err)
	}
	sound := processor.GetSound()
	fmt.Println(sound)
}
