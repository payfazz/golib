package writer

import (
	"github.com/payfazz/golib/pkg/slack/client"
)

// Writer ...
type Writer struct {
	author string
	title  string
	color  string
	footer string
	slack  *client.Client
}

func (writer *Writer) Write(p []byte) (int, error) {
	msg := string(p)
	len := len(msg)
	err := writer.slack.Send(writer.author, writer.title, writer.footer, writer.color, msg)
	return len, err
}

// New ...
func New(webhookURI, author, title, color, footer string) *Writer {
	slack := client.New(webhookURI)
	return &Writer{
		author: author,
		title:  title,
		color:  color,
		footer: footer,
		slack:  slack,
	}
}
