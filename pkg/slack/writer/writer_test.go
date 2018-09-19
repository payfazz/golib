package writer_test

import (
	"log"
	"testing"

	"github.com/payfazz/golib/pkg/slack/writer"
)

const webhookURI string = "https://hooks.slack.com/services/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"

func TestWriterWithString(t *testing.T) {
	writer, errWriter := writer.New(webhookURI, "writer_test", "TestWriter", "#dddddd", "")
	if errWriter != nil {
		t.Error(errWriter)
	}
	logger := log.New(writer, "", log.Llongfile)
	logger.Print("hello logger")
}

type testStruct struct {
	Number int      `json:"number"`
	Text   string   `json:"text"`
	Array  []string `json:"array"`
}

func TestWriterWithStruct(t *testing.T) {
	writer, errWriter := writer.New(webhookURI, "writer_test", "TestWriter", "#dddddd", "")
	if errWriter != nil {
		t.Error(errWriter)
	}
	logger := log.New(writer, "", log.LstdFlags)
	data := &testStruct{}
	logger.Printf("%+v", data)
}
