package slack

import (
	"encoding/json"
	"fmt"
	"time"
)

// Message ...
type Message struct {
	Text        string        `json:"text,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// Attachment ...
type Attachment struct {
	Fallback     string   `json:"fallback"`
	Color        string   `json:"color"`
	AuthorName   string   `json:"author_name"`
	Title        string   `json:"title"`
	Text         string   `json:"text"`
	ThumbnailURI string   `json:"thumb_url,omitempty"`
	MarkDownIn   []string `json:"mrkdwn_in,omitempty"`
	Footer       string   `json:"footer,omitempty"`
	FooterIcon   string   `json:"footer_icon,omitempty"`
	Timestamp    int64    `json:"ts"`
}

// NewDefaultAttachment ...
func (m *Message) NewDefaultAttachment(author, title, footer string, data interface{}) (*Attachment, error) {
	return m.NewAttachment(author, title, footer, "#006ce0", data)
}

// NewSuccessAttachment ...
func (m *Message) NewSuccessAttachment(author, title, footer string, data interface{}) (*Attachment, error) {
	return m.NewAttachment(author, title, footer, "#71c400", data)
}

// NewErrorAttachment ...
func (m *Message) NewErrorAttachment(author, title, footer string, data interface{}) (*Attachment, error) {
	return m.NewAttachment(author, title, footer, "#ff0048", data)
}

// NewAttachment ...
func (m *Message) NewAttachment(author, title, footer, color string, data interface{}) (*Attachment, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	attachment := &Attachment{
		AuthorName: author,
		Color:      color,
		Title:      title,
		Text:       fmt.Sprintf("```%s```", string(jsonBytes)),
		Footer:     footer,
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",
		MarkDownIn: []string{"text"},
	}
	attachment.Timestamp = time.Now().Unix()
	m.Attachments = append(m.Attachments, attachment)
	return attachment, nil
}

// NewMessage ...
func NewMessage() *Message {
	m := &Message{
		Attachments: []*Attachment{},
	}
	return m
}
