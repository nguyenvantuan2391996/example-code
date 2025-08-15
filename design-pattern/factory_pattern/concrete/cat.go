package concrete

type Cat struct {
}

func NewCat() *Cat {
	return &Cat{}
}

func (c *Cat) GetSound() string {
	return "meow meow"
}
