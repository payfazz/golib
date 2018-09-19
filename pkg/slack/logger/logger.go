package logger

import (
	"log"

	"github.com/payfazz/golib/pkg/slack/writer"
)

// New slack logger
func New(webhookURI, author, title, footer, color string) *log.Logger {
	writer := writer.New(webhookURI, author, title, color, footer)
	logger := log.New(writer, "", 0)
	return logger
}
