package teams

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var webhookTeamsClient *WebhookTeamsClient

type WebhookTeamsClient struct {
	client     *http.Client
	webhookURL string
}

func NewWebhookTeamsClient(webhookURL string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	webhookTeamsClient = &WebhookTeamsClient{
		client: &http.Client{
			Timeout:   20 * time.Second,
			Transport: tr,
		},
		webhookURL: webhookURL,
	}

	logrus.Info("completed initializing incoming webhook teams")
}

func GetInstance() *WebhookTeamsClient {
	return webhookTeamsClient
}

func (w *WebhookTeamsClient) BuildBody(texts []string) []*Body {
	body := []*Body{
		{
			Type: TypeImage,
			URL:  ActivityImage,
			Size: SizeMedium,
		},
	}

	for _, text := range texts {
		body = append(body, &Body{
			Type: TypeTextBlock,
			Text: text,
			Wrap: true,
		})
	}

	return body
}

func (w *WebhookTeamsClient) Send(templateMsg *TemplateTeamsMessage) error {
	payload, err := json.Marshal(templateMsg)
	if err != nil {
		return err
	}

	webhookTeamsReq, err := http.NewRequest(
		http.MethodPost,
		w.webhookURL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	webhookTeamsReq.Header.Add("Content-Type", "application/json")

	resp, err := w.client.Do(webhookTeamsReq)
	if resp != nil {
		defer func(body io.ReadCloser) {
			CloseResponse(body)
		}(resp.Body)
	}

	if err != nil || resp == nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed send teams message")
	}

	logrus.Info(fmt.Sprintf("send teams message is completed | status code: %d", resp.StatusCode))

	return nil
}

func (w *WebhookTeamsClient) SendMessage(templateMsg *TemplateTeamsMessage) error {
	if len(w.webhookURL) == 0 {
		return errors.New("webhook url is empty")
	}

	var err error
	for i := 0; i < MaxRetryTimes; i++ {
		err = w.Send(templateMsg)
		if err == nil {
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	return err
}

func (w *WebhookTeamsClient) SendV2(templateMsg *TemplateTeamsMessageV2) error {
	payload, err := json.Marshal(templateMsg)
	if err != nil {
		return err
	}

	webhookTeamsReq, err := http.NewRequest(
		http.MethodPost,
		w.webhookURL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	webhookTeamsReq.Header.Add("Content-Type", "application/json")

	resp, err := w.client.Do(webhookTeamsReq)
	if resp != nil {
		defer func(body io.ReadCloser) {
			CloseResponse(body)
		}(resp.Body)
	}

	if err != nil || resp == nil {
		return err
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted) {
		return fmt.Errorf("failed send teams message")
	}

	logrus.Info(fmt.Sprintf("send teams message is completed | status code: %d", resp.StatusCode))

	return nil
}

func (w *WebhookTeamsClient) SendMessageV2(body []*Body, summary string) error {
	if len(w.webhookURL) == 0 {
		return fmt.Errorf("webhook url is empty")
	}

	content := &Content{
		Type: TypeAdaptiveCard,
		Body: body,
	}

	attachments := []*Attachment{
		{
			ContentType: ContentTypeAdaptiveCard,
			Content:     content,
		},
	}

	templateMessageV2 := &TemplateTeamsMessageV2{
		Type:        TypeMessage,
		Summary:     summary,
		Attachments: attachments,
	}

	var err error
	for i := 0; i < MaxRetryTimes; i++ {
		err = w.SendV2(templateMessageV2)
		if err == nil {
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	return err
}
