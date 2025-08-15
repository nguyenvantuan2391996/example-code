package main

import (
	"fmt"

	"example-code/design-pattern/builder_pattern/builder"
	"example-code/design-pattern/builder_pattern/constant"
	"example-code/design-pattern/builder_pattern/director"
	"example-code/design-pattern/builder_pattern/product"
)

func main() {
	in := &product.InfoInput{
		Price:    100,
		Discount: 10,
		Amount:   15,
	}

	builderApple := builder.GetFruitBuilder(constant.Apple)
	directorApple := director.NewDirector(builderApple)
	outApple := directorApple.BuildOutput(in)

	builderOrange := builder.GetFruitBuilder(constant.Orange)
	directorOrange := director.NewDirector(builderOrange)
	outOrange := directorOrange.BuildOutput(in)

	fmt.Printf("Pay for orange : $%v\n", outOrange.MoneyPayment)
	fmt.Printf("Pay for apple : $%v\n", outApple.MoneyPayment)
}
