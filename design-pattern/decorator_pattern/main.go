package main

import (
	"fmt"
	"time"

	concreteComponent "example-code/design-pattern/decorator_pattern/concrete_component"
	concreteDecorator "example-code/design-pattern/decorator_pattern/concrete_decorator"
)

func main() {
	// If you create your account then we will send the notification to your SMS
	CreateAccount("Tuan Nguyen")

	// If you update your account, eg: password, name,... then we will send the notification to your SMS, your Facebook and our Slack
	UpdateAccount("Tuan Nguyen")

	// If you delete your account then we will send the notification to your SMS, our Slack
	DeleteAccount("Tuan Nguyen")
}

func CreateAccount(name string) {
	fmt.Println(fmt.Sprintf("Create the account with name %s is done.", name))
	time.Sleep(2 * time.Second)

	fmt.Println("Begin to send the notification...")
	time.Sleep(2 * time.Second)

	notifier := &concreteComponent.Notifier{}

	// add sms notification
	sms := &concreteDecorator.SmsNotifier{
		Notifier: notifier,
	}

	// send the notification
	sms.SendNotification("create account")
}

func UpdateAccount(name string) {
	fmt.Println(fmt.Sprintf("Update the account with name %s is done.", name))
	time.Sleep(2 * time.Second)

	fmt.Println("Begin to send the notification...")
	time.Sleep(2 * time.Second)

	notifier := &concreteComponent.Notifier{}

	// add sms notification
	sms := &concreteDecorator.SmsNotifier{Notifier: notifier}

	// add facebook notification
	facebookSMS := &concreteDecorator.FacebookNotifier{Notifier: sms}

	// add Slack notification
	slackFacebookSMS := &concreteDecorator.SlackNotifier{Notifier: facebookSMS}

	// send the notification
	slackFacebookSMS.SendNotification("update account")
}

func DeleteAccount(name string) {
	fmt.Println(fmt.Sprintf("Delete the account with name %s is done.", name))
	time.Sleep(2 * time.Second)

	fmt.Println("Begin to send the notification...")
	time.Sleep(2 * time.Second)

	notifier := &concreteComponent.Notifier{}

	// add sms notification
	sms := &concreteDecorator.SmsNotifier{Notifier: notifier}

	// add Slack notification
	slackSMS := &concreteDecorator.SlackNotifier{Notifier: sms}

	// send the notification
	slackSMS.SendNotification("delete account")
}
