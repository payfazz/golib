package slack

import (
	e "github.com/payfazz/golib/pkg/errors"

	"github.com/payfazz/golib/pkg/slack"
)

// Logger ...
type Logger struct {
	client *slack.Client
}

// Log ...
func (l *Logger) Log(source, event string, data interface{}) error {
	return l.client.Send(source, event, source, "#dddddd", data)
}

// LogWarning ...
func (l *Logger) LogWarning(source, event string, data interface{}) error {
	return l.client.Send(source, event, source, "#ffc700", data)
}

// LogError ...
func (l *Logger) LogError(source, event string, data interface{}) error {
	return l.client.Send(source, event, source, "#ff0048", data)
}

// NewLogger ...
func NewLogger(hookURI string) (*Logger, error) {
	slackClient, errClient := slack.NewClient(hookURI)
	if errClient != nil {
		return nil, e.ServiceError("slack.logger", errClient)
	}
	l := &Logger{
		client: slackClient,
	}
	return l, nil
}
