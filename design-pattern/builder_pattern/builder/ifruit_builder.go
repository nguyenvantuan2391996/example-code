package builder

import (
	"example-code/design-pattern/builder_pattern/concreate_builder"
	"example-code/design-pattern/builder_pattern/constant"
	"example-code/design-pattern/builder_pattern/product"
)

type IFruitBuilder interface {
	SetPrice(in *product.InfoInput)
	SetDiscount(in *product.InfoInput)
	SetMoneyPayment(in *product.InfoInput)
	ToOutputMoney() *product.InfoOutput
}

func GetFruitBuilder(fruitName string) IFruitBuilder {
	switch fruitName {
	case constant.Orange:
		return concreate_builder.NewOrangeBuilder()
	case constant.Apple:
		return concreate_builder.NewAppleBuilder()
	}
	return nil
}
