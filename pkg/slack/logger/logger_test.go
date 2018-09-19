package logger_test

import (
	"testing"

	slack "github.com/payfazz/golib/pkg/slack/logger"
)

const webhookURI string = "https://hooks.slack.com/services/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"

func TestLogSimpleString(t *testing.T) {
	l := slack.New(webhookURI, "slack/logger", "TestLogSimpleString", "", "#dddddd")
	l.Print("hello, TestLogSimpleString")
}
