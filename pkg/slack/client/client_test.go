package client

import (
	"testing"
)

const webhookURI string = "https://hooks.slack.com/services/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"

var client *Client

func init() {
	client = New(webhookURI)
}

func TestText(t *testing.T) {
	errText := client.Text("Hello text")
	if errText != nil {
		t.Error(errText)
	}
}

func TestAttachment(t *testing.T) {
	att := NewAttachment()
	// att.Title = "Hello attachment"
	att.Text = "Hello attachment"
	// att.Footer = "Hello attachment"
	errText := client.Attachment(att)
	if errText != nil {
		t.Error(errText)
	}
}
