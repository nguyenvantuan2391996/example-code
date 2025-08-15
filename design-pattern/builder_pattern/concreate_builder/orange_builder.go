package concreate_builder

import "example-code/design-pattern/builder_pattern/product"

type OrangeBuilder struct {
	Price        float64
	Discount     float64
	MoneyPayment float64
}

func NewOrangeBuilder() *OrangeBuilder {
	return &OrangeBuilder{}
}

func (o *OrangeBuilder) SetPrice(in *product.InfoInput) {
	if in.Amount > 10 {
		o.Price = in.Price - 5
	}
}

func (o *OrangeBuilder) SetDiscount(in *product.InfoInput) {
	if in.Amount > 10 {
		o.Discount = in.Discount + 2
	}
}

func (o *OrangeBuilder) SetMoneyPayment(in *product.InfoInput) {
	o.MoneyPayment = o.Price * (100 - o.Discount/100) * float64(in.Amount)
}

func (o *OrangeBuilder) ToOutputMoney() *product.InfoOutput {
	return &product.InfoOutput{
		MoneyPayment: o.MoneyPayment,
	}
}
