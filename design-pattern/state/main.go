package main

import "fmt"

type State interface {
    InsertCoin()
    Dispense()
}

type VendingMachine struct {
    state State
}

func (v *VendingMachine) SetState(s State) {
    v.state = s
}

func (v *VendingMachine) InsertCoin() { v.state.InsertCoin() }
func (v *VendingMachine) Dispense()   { v.state.Dispense() }

type NoCoinState struct {
    machine *VendingMachine
}

func (s NoCoinState) InsertCoin() {
    fmt.Println("Coin inserted → Ready to dispense")
    s.machine.SetState(HaveCoinState{s.machine})
}

func (NoCoinState) Dispense() {
    fmt.Println("Insert coin first")
}

type HaveCoinState struct {
    machine *VendingMachine
}

func (s HaveCoinState) InsertCoin() {
    fmt.Println("Already have a coin")
}

func (s HaveCoinState) Dispense() {
    fmt.Println("Item dispensed")
    s.machine.SetState(NoCoinState{s.machine})
}

func main() {
    vm := &VendingMachine{}
    vm.SetState(NoCoinState{vm})

    vm.Dispense()
    vm.InsertCoin()
    vm.Dispense()
}