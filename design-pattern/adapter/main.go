package main

import "fmt"

type Payment interface {
    Pay(amount int)
}

type MomoSDK struct{}

func (MomoSDK) SendMoney(v int) {
    fmt.Printf("Momo paid %d VND\n", v)
}

type MomoAdapter struct {
    momo MomoSDK
}

func (a MomoAdapter) Pay(amount int) {
    a.momo.SendMoney(amount)
}

func main() {
    var p Payment = MomoAdapter{momo: MomoSDK{}}
    p.Pay(50000)
}