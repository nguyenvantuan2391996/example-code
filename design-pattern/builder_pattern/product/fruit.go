package product

type InfoInput struct {
	Price    float64
	Discount float64
	Amount   int64
}

type InfoOutput struct {
	MoneyPayment float64
}
