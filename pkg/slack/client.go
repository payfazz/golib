package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	e "github.com/payfazz/golib/pkg/errors"
)

// SendText ...
func (c *Client) SendText(text string) error {
	message := NewMessage()
	message.Text = text
	return c.SendMessage(message)
}

// Send ...
func (c *Client) Send(author, title, footer, color string, jsonConvertible interface{}) error {
	message := NewMessage()
	_, err := message.NewAttachment(author, title, footer, color, jsonConvertible)
	if err != nil {
		return e.ServiceError("slack", err)
	}
	return c.SendMessage(message)
}

// SendAttachment ...
func (c *Client) SendAttachment(att *Attachment) error {
	message := NewMessage()
	message.Attachments = append(message.Attachments, att)
	return c.SendMessage(message)
}

// SendMessage ...
func (c *Client) SendMessage(message *Message) error {
	msgBytes, errMsgMarshal := json.Marshal(message)
	if errMsgMarshal != nil {
		return e.ServiceError("slack", errMsgMarshal)
	}
	request, errReq := http.NewRequest("POST", c.hookURI, bytes.NewBuffer(msgBytes))
	if errReq != nil {
		return e.ServiceError("slack", errReq)
	}
	request.Header.Set("Content-Type", "application/json")
	response, errResponse := c.httpClient.Do(request)
	if errResponse != nil {
		return e.ServiceError("slack", errResponse)
	}
	if response.StatusCode != http.StatusOK {
		return e.ServiceError("slack", errors.New("failed to send message to slack"))
	}
	return nil
}

// Client ...
type Client struct {
	hookURI    string
	httpClient *http.Client
}

// NewClient ...
func NewClient(hookURI string) (*Client, error) {
	c := &Client{
		hookURI:    hookURI,
		httpClient: &http.Client{},
	}
	return c, nil
}
