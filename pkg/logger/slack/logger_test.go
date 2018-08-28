package slack_test

import (
	"testing"

	"github.com/payfazz/golib/pkg/logger/slack"
)

func TestLogSimpleString(t *testing.T) {
	l, e := slack.NewLogger("https://hooks.slack.com/services/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")
	if e != nil {
		t.Error(e)
	}
	l.Log("github.com/payfazz/golib/pkg/logger/slack/logger_test.go", "TestLogSimpleString", "test")
}

func TestLogWarningSimpleString(t *testing.T) {
	l, e := slack.NewLogger("https://hooks.slack.com/services/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")
	if e != nil {
		t.Error(e)
	}
	l.LogWarning("github.com/payfazz/golib/pkg/logger/slack/logger_test.go", "TestLogWarningSimpleString", "test")
}

func TestLogErrorSimpleString(t *testing.T) {
	l, e := slack.NewLogger("https://hooks.slack.com/services/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")
	if e != nil {
		t.Error(e)
	}
	l.LogError("github.com/payfazz/golib/pkg/logger/slack/logger_test.go", "TestLogErrorSimpleString", "test")
}
