package main

import "fmt"

type Mediator interface {
    Send(msg string, user *User)
}

type ChatRoom struct{}

func (ChatRoom) Send(msg string, user *User) {
    fmt.Printf("[%s]: %s\n", user.name, msg)
}

type User struct {
    name     string
    mediator Mediator
}

func (u User) Send(msg string) {
    u.mediator.Send(msg, &u)
}

func main() {
    room := ChatRoom{}
    u1 := User{"Ton", room}
    u2 := User{"Duc", room}

    u1.Send("Hello!")
    u2.Send("Hi bro!")
}