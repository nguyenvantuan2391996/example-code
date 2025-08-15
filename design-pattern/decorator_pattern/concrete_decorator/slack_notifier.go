package concrete_decorator

import (
	"fmt"
	"time"

	"example-code/design-pattern/decorator_pattern/component"
)

type SlackNotifier struct {
	Notifier component.INotifier
}

func (s *SlackNotifier) SendNotification(msg string) {
	s.Notifier.SendNotification(msg)
	fmt.Println(fmt.Sprintf("Send the notification to Slack with message %s", msg))
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
