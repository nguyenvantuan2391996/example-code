package teams

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	webhookURL = ""
)

func TestWebhookTeamsClient_SendMessage(t *testing.T) {
	viper.Set("WEBHOOK_URL", webhookURL)

	NewWebhookTeamsClient(viper.GetString("WEBHOOK_URL"))

	instance := GetInstance()

	facts := []*Fact{
		{
			Name:  "Message",
			Value: "Test",
		},
	}

	sections := []*Section{
		{
			ActivityTitle:    "Send message",
			ActivitySubtitle: "HCM",
			ActivityImage:    "https://adaptivecards.io/content/cats/3.png",
			Facts:            facts,
			Markdown:         true,
		},
	}

	err := instance.SendMessage(&TemplateTeamsMessage{
		Type:     TypeMessageCard,
		Summary:  "Notify",
		Sections: sections,
	})
	assert.NoError(t, err)
}

func TestWebhookTeamsClient_SendMessageV2(t *testing.T) {
	viper.Set("WEBHOOK_URL", webhookURL)

	NewWebhookTeamsClient(viper.GetString("WEBHOOK_URL"))

	instance := GetInstance()

	err := instance.SendMessageV2(instance.BuildBody([]string{
		"**ACM Name:** Solare Restaurant",
		"**Latency:** 37 seconds",
	}), "Notify")

	assert.NoError(t, err)
}
