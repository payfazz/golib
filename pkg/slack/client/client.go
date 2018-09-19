package client

import (
	"bytes"
	"encoding/json"
	e "errors"
	"net/http"
	"time"

	"github.com/payfazz/golib/pkg/errors"
)

// Text ...
func (client *Client) Text(text string) error {
	message := &Message{
		Text: text,
	}
	return client.Message(message)
}

// Attachment ...
func (client *Client) Attachment(attachment *Attachment) error {
	message := NewMessage()
	message.Attachments.Add(attachment)
	return client.Message(message)
}

// Send ...
func (client *Client) Send(author, title, footer, color string, jsonConvertible interface{}) error {
	message := NewMessage()
	_, err := message.Add(author, title, footer, color, jsonConvertible)
	if err != nil {
		return errors.ServiceError("slack", err)
	}
	return client.Message(message)
}

// Message ...
func (client *Client) Message(message *Message) error {
	if message == nil {
		return errors.ServiceError("slack", e.New("message_required"))
	}
	if message.Attachments != nil {
		now := time.Now().UTC()
		for _, attachment := range []*Attachment(*message.Attachments) {
			attachment.Timestamp = now.Unix()
		}
	}
	msgBytes, errMsgMarshal := json.Marshal(message)
	if errMsgMarshal != nil {
		return errors.ServiceError("slack", errMsgMarshal)
	}
	request, errReq := http.NewRequest("POST", client.hookURI, bytes.NewBuffer(msgBytes))
	if errReq != nil {
		return errors.ServiceError("slack", errReq)
	}
	request.Header.Set("Content-Type", "application/json")
	response, errResponse := client.httpClient.Do(request)
	if errResponse != nil {
		return errors.ServiceError("slack", errResponse)
	}
	if response.StatusCode != http.StatusOK {
		return errors.ServiceError("slack", e.New("failed to send message to slack"))
	}
	return nil
}

// Client ...
type Client struct {
	hookURI    string
	httpClient *http.Client
}

// New ...
func New(hookURI string) *Client {
	client := &Client{
		hookURI:    hookURI,
		httpClient: &http.Client{},
	}
	return client
}
