package component

type INotifier interface {
	SendNotification(msg string)
}
