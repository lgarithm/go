package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Webhook is https://api.slack.com/incoming-webhooks
type Webhook struct {
	Endpoint string
}

// Report implements https://api.slack.com/incoming-webhooks#sending_messages
func (w Webhook) Report(text string) error {
	msg := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(msg)
	client := http.Client{}
	res, err := client.Post(w.Endpoint, "application/json", &b)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}
