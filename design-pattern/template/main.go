package main

import "fmt"

type Cooker interface {
    Prepare()
    Cook()
    Serve()
}

func MakeDish(c Cooker) {
    c.Prepare()
    c.Cook()
    c.Serve()
}

type Pho struct{}

func (Pho) Prepare() { fmt.Println("Prepare noodles & broth") }
func (Pho) Cook()    { fmt.Println("Cooking Pho") }
func (Pho) Serve()   { fmt.Println("Serving Pho") }

type BunBo struct{}

func (BunBo) Prepare() { fmt.Println("Prepare meat & spices") }
func (BunBo) Cook()    { fmt.Println("Cooking Bun Bo") }
func (BunBo) Serve()   { fmt.Println("Serving Bun Bo Hue") }

func main() {
    MakeDish(Pho{})
    MakeDish(BunBo{})
}