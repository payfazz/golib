package client

import (
	"encoding/json"
	"fmt"
)

// Message ...
type Message struct {
	Text        string       `json:"text,omitempty"`
	Attachments *Attachments `json:"attachments,omitempty"`
}

// Add ...
func (message *Message) Add(author, title, footer, color string, data interface{}) (*Attachment, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	attachment := NewAttachment()
	attachment.AuthorName = author
	attachment.Color = color
	attachment.Title = title
	attachment.Text = fmt.Sprintf("```%s```", string(jsonBytes))
	attachment.Footer = footer
	message.AddAttachment(attachment)
	return attachment, nil
}

// AddAttachment ...
func (message *Message) AddAttachment(attachment *Attachment) {
	message.Attachments.Add(attachment)
}

// NewMessage ...
func NewMessage() *Message {
	message := &Message{
		Attachments: &Attachments{},
	}
	return message
}
