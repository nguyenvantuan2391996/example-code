package concrete_decorator

import (
	"fmt"
	"time"

	"example-code/design-pattern/decorator_pattern/component"
)

type FacebookNotifier struct {
	Notifier component.INotifier
}

func (f *FacebookNotifier) SendNotification(msg string) {
	f.Notifier.SendNotification(msg)
	fmt.Println(fmt.Sprintf("Send the notification to Facebook with message %s", msg))
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
