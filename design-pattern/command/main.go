package main

import "fmt"

// Receiver
type Light struct{}

func (Light) On()  { fmt.Println("Light ON") }
func (Light) Off() { fmt.Println("Light OFF") }

// Command interface
type Command interface {
    Execute()
}

// Command A
type OnCommand struct {
    light Light
}

func (c OnCommand) Execute() { c.light.On() }

// Command B
type OffCommand struct {
    light Light
}

func (c OffCommand) Execute() { c.light.Off() }

// Invoker
type Button struct {
    command Command
}

func (b Button) Press() { b.command.Execute() }

func main() {
    light := Light{}
    onBtn := Button{command: OnCommand{light}}
    offBtn := Button{command: OffCommand{light}}

    onBtn.Press()
    offBtn.Press()
}