package concrete

type Rat struct {
}

func NewRat() *Rat {
	return &Rat{}
}

func (r *Rat) GetSound() string {
	return "chit chit"
}
