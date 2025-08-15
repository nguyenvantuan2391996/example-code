package director

import (
	"example-code/design-pattern/builder_pattern/builder"
	"example-code/design-pattern/builder_pattern/product"
)

type Director struct {
	FruitBuilder builder.IFruitBuilder
}

func NewDirector(f builder.IFruitBuilder) *Director {
	return &Director{
		FruitBuilder: f,
	}
}

func (d *Director) BuildOutput(in *product.InfoInput) *product.InfoOutput {
	d.FruitBuilder.SetPrice(in)
	d.FruitBuilder.SetDiscount(in)
	d.FruitBuilder.SetMoneyPayment(in)
	return d.FruitBuilder.ToOutputMoney()
}
