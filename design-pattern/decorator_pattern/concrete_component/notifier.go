package concrete_component

import (
	"fmt"
	"time"
)

type Notifier struct {
}

func (n *Notifier) SendNotification(msg string) {
	fmt.Println("Handle process logic common...")
	time.Sleep(2 * time.Second)
}
