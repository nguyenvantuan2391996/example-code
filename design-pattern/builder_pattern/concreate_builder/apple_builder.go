package concreate_builder

import "example-code/design-pattern/builder_pattern/product"

type AppleBuilder struct {
	Price        float64
	Discount     float64
	MoneyPayment float64
}

func NewAppleBuilder() *AppleBuilder {
	return &AppleBuilder{}
}

func (o *AppleBuilder) SetPrice(in *product.InfoInput) {
	if in.Amount > 10 {
		o.Price = in.Price - 10
	}
}

func (o *AppleBuilder) SetDiscount(in *product.InfoInput) {
	if in.Amount > 10 {
		o.Discount = in.Discount + 5
	}
}

func (o *AppleBuilder) SetMoneyPayment(in *product.InfoInput) {
	o.MoneyPayment = o.Price * (100 - o.Discount/100) * float64(in.Amount)
}

func (o *AppleBuilder) ToOutputMoney() *product.InfoOutput {
	return &product.InfoOutput{
		MoneyPayment: o.MoneyPayment,
	}
}
