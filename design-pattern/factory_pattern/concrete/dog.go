package concrete

type Dog struct {
}

func NewDog() *Dog {
	return &Dog{}
}

func (d *Dog) GetSound() string {
	return "gauw gauw"
}
