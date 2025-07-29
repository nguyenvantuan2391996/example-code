package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	DefaultEmoji = ":ghost:"
)

func sendMessage(webhook string, message *SlackMessage) error {
	logrus.Info("Send message to webhook ", webhook)

	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", webhook, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		if res != nil {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				logrus.Warnf("response %v", string(bodyBytes))
				return err
			}
			return errors.New("notification failure")
		}
	}
	logrus.Info("successfully sent notification...")
	return nil
}

func SendMessageSlack(webhook string, slackMessage *SlackMessage) (err error) {
	if len(webhook) == 0 {
		return errors.New("missing channel url")
	}

	for i := 0; i < 5; i++ {
		err = sendMessage(webhook, slackMessage)
		if err == nil {
			return
		}
	}
	return
}