package client

import (
	"testing"
)

// TestAddAttachment ...
func TestAddAttachment(t *testing.T) {
	attachments := Attachments{}
	attachments.Add(NewAttachment())
	n := len(attachments)
	if n != 1 {
		t.Error("expected len 1")
	}
}
